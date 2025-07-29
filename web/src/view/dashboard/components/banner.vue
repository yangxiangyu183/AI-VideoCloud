<template>
  <el-carousel
    :interval="4000"
    :arrow="'hover'"
    :height="'100%'"
    indicator-position="none"
  >
    <el-carousel-item
      class="cursor-pointer lg:h-40"
      v-for="(item, index) in banners"
      :key="index"
      @click="openLink(item.link)"
    >
      <img class="w-full h-full object-cover" :src="item.img" alt="" />
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
  import banner from '@/assets/banner.jpg'
  import banner2 from '@/assets/banner2.jpg'

  const openLink = (link) => {
    window.open(link)
  }

  const banners = [
    {
      img: banner,
      link: 'https://gin-vue-admin.com/empower/index.html'
    },
    {
      img: banner2,
      link: 'https://plugin.gin-vue-admin.com'
    },
  ]
</script>

<style scoped lang="scss"></style>