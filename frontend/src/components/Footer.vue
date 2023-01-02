<template>
  <div class="container py-4">
    <div class="row">
      <div class="col-md"></div>
      <div class="col-md-4 border-bottom">
        <p style="font-weight: bold">Sosyal Medya</p>
      </div>
      <div class="col-md"></div>
      <div class="col-md-4 border-bottom">
        <p>Deneme2</p>
      </div>
      <div class="col-md"></div>
    </div>
    <div class="row">
      <div class="col-md"></div>
      <div class="col-md-4">
        <img src="../assets/image/Blog-logo.png" width="45"/>
      </div>
      <div class="col-md"></div>
      <div class="col-md-4">
        <p class="copyright">Hbostan © Copyright 2022</p>
      </div>
      <div class="col-md"></div>
    </div>
  </div>
</template>

<!--
* Col row ayarlandı
* border çizgisi çekildi
* footer
-->

<script>
export default {
  name: "appFooter"
}
</script>

<style scoped>

.container {
  background-color: white;
}

.copyright{
  text-align: end;
}

</style>