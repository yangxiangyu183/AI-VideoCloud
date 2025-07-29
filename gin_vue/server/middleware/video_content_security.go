package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

// 配置腾讯云API密钥和COS信息
var (
	SecretId      = global.GVA_CONFIG.TencentCOS.SecretID  // 从配置文件读取
	SecretKey     = global.GVA_CONFIG.TencentCOS.SecretKey // 从配置文件读取
	Region        = "ap-shanghai"                          // 地域，如上海：ap-shanghai
	CMS_Endpoint  = "cms.tencentcloudapi.com"
	CMS_Version   = "2020-12-29"
	COS_Bucket    = global.GVA_CONFIG.TencentCOS.Bucket // 从配置文件读取
	COS_Region    = "ap-shanghai"                       // COS地域
	COS_KeyPrefix = "videos/"                           // 视频在COS中的存储路径前缀
)

// 上传本地视频到腾讯云COS
func UploadVideoToCOS(localFilePath string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("本地文件不存在: %s", localFilePath)
	}

	// 解析文件信息
	fileName := filepath.Base(localFilePath)
	key := fmt.Sprintf("%s%s_%d%s", COS_KeyPrefix,
		strings.TrimSuffix(fileName, filepath.Ext(fileName)),
		time.Now().Unix(), filepath.Ext(fileName))

	// 初始化COS客户端
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", COS_Bucket, COS_Region))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretId,
			SecretKey: SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	// 打开本地文件
	file, err := os.Open(localFilePath)
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 获取文件信息
	_, err = file.Stat()
	if err != nil {
		return "", fmt.Errorf("获取文件信息失败: %v", err)
	}

	// 上传文件到COS
	_, err = client.Object.PutFromFile(context.Background(), key, localFilePath, nil)
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %v", err)
	}

	// 生成文件的访问URL
	fileURL := fmt.Sprintf("https://%s.cos.%s.myqcloud.com/%s", COS_Bucket, COS_Region, key)
	fmt.Printf("文件上传成功，URL: %s\n", fileURL)

	return fileURL, nil
}

// 提交视频审核任务
func SubmitVideoModerationTask(videoURL string) (string, error) {
	// 构建请求参数
	params := map[string]interface{}{
		"Action":  "CreateVideoModerationTask",
		"Version": "2020-12-29",
		"Region":  "ap-shanghai",
		"Input": map[string]interface{}{
			"Url": videoURL,
		},
		"Conf": map[string]interface{}{
			"DetectType": []string{"Porn", "Terrorism", "Politics", "Ad", "Illegal", "Abuse"},
		},
	}

	// 发送请求
	respData, err := sendTencentCloudRequest(params, "cms.tencentcloudapi.com")
	if err != nil {
		return "", err
	}

	// 解析响应
	result := make(map[string]interface{})
	if err := json.Unmarshal(respData, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查是否有错误
	if _, ok := result["Error"]; ok {
		return "", fmt.Errorf("API错误: %v", result["Error"])
	}

	// 提取任务ID
	response, ok := result["Response"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("无效的响应格式")
	}

	taskId, ok := response["TaskId"].(string)
	if !ok {
		return "", fmt.Errorf("无法获取任务ID")
	}

	return taskId, nil
}

// 查询视频审核结果
func QueryModerationResult(taskId string) (map[string]interface{}, error) {
	// 构建请求参数
	params := map[string]interface{}{
		"Action":  "DescribeVideoModerationTask",
		"Version": "2020-12-29",
		"Region":  "ap-shanghai",
		"TaskId":  taskId,
	}

	// 发送请求
	respData, err := sendTencentCloudRequest(params, "cms.tencentcloudapi.com")
	if err != nil {
		return nil, err
	}

	// 解析响应
	result := make(map[string]interface{})
	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查是否有错误
	if _, ok := result["Error"]; ok {
		return nil, fmt.Errorf("API错误: %v", result["Error"])
	}

	return result, nil
}

// 打印审核结果
func PrintModerationResult(result map[string]interface{}) {
	response, ok := result["Response"].(map[string]interface{})
	if !ok {
		fmt.Println("无效的结果格式")
		return
	}

	fmt.Println("视频审核结果:")
	fmt.Printf("任务ID: %s\n", response["TaskId"])
	fmt.Printf("任务状态: %s\n", response["Status"])

	// 检查是否有审核建议
	if suggestion, ok := response["Suggestion"].(string); ok {
		fmt.Printf("审核建议: %s\n", suggestion)
	}

	// 检查是否有审核结论
	if label, ok := response["Label"].(string); ok {
		fmt.Printf("审核结论: %s\n", label)
	}

	// 打印详细审核结果
	if results, ok := response["Results"].([]interface{}); ok {
		fmt.Println("\n详细审核结果:")
		for i, res := range results {
			resultItem, ok := res.(map[string]interface{})
			if !ok {
				continue
			}
			fmt.Printf("结果 %d:\n", i+1)
			if scene, ok := resultItem["Scene"].(string); ok {
				fmt.Printf("  场景: %s\n", scene)
			}
			if label, ok := resultItem["Label"].(string); ok {
				fmt.Printf("  标签: %s\n", label)
			}
			if score, ok := resultItem["Score"].(float64); ok {
				fmt.Printf("  得分: %.2f\n", score)
			}
			if suggestion, ok := resultItem["Suggestion"].(string); ok {
				fmt.Printf("  建议: %s\n", suggestion)
			}
		}
	}
}

// 发送腾讯云API请求
func sendTencentCloudRequest(params map[string]interface{}, endpoint string) ([]byte, error) {
	// 添加公共参数
	params["Nonce"] = time.Now().UnixNano() / 1000000
	params["Timestamp"] = time.Now().Unix()
	params["SecretId"] = SecretId

	// 生成签名
	signature := generateSignature(params, endpoint)

	// 构建请求URL
	urlParams := fmt.Sprintf("Signature=%s", url.QueryEscape(signature))
	for key, value := range params {
		urlParams += fmt.Sprintf("&%s=%s", key, encodeValue(value))
	}

	url := fmt.Sprintf("https://%s?%s", endpoint, urlParams)

	// 发送请求
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败，状态码: %d，响应: %s", resp.StatusCode, string(respData))
	}

	return respData, nil
}

// 生成腾讯云API签名
func generateSignature(params map[string]interface{}, endpoint string) string {
	// 提取并排序参数键
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 构建待签名字符串
	var signStr string
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, encodeValue(params[key]))
	}
	signStr = "POST" + endpoint + "/" + "?" + strings.TrimRight(signStr, "&")

	// 计算HMAC-SHA1签名
	mac := hmac.New(sha1.New, []byte(SecretKey))
	mac.Write([]byte(signStr))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature
}

// 编码参数值
func encodeValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return url.QueryEscape(v)
	case int64:
		return fmt.Sprintf("%d", v)
	case []string:
		return url.QueryEscape(strings.Join(v, ","))
	default:
		jsonStr, _ := json.Marshal(v)
		return url.QueryEscape(string(jsonStr))
	}
}
