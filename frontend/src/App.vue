<template>
  <div class="container">
    <app-header></app-header>
    <app-contents></app-contents>
    <app-footer></app-footer>
  </div>
</template>

<script>
import Header from "./components/Header";
import Contents from "./components/Contents";
import Footer from "./components/Footer";

export default {
  name: 'App',
  components: {
    appHeader : Header,
    appContents : Contents,
    appFooter : Footer
  }
}
</script>

<style>
.body {
  background-color: #f5f5f5;
}
</style>
