<template>
  <div class="row">
    <header
      class="d-flex  align-items-center justify-content-center justify-content-md-between py-3 mb-3 border-bottom">
      <div class="d-flex align-items-center col-md-3 mb-md-0 img-fluid">
        <img src="../assets/image/Blog-logo.png" width="45"/>
      </div>
      <ul class="nav col-12 col-md-auto mb-2 justify content-center mb-md-0">
        <li>
          <a href="../pages/MainPage" class="nav-link" >Ana Sayfa</a>
        </li>
        <li>
          <a href="../pages/Categories" class="nav-link">Kategoriler</a>
        </li>
        <li>
          <a href="../pages/Articles" class="nav-link">Makaleler</a>
        </li>
        <li>
          <a href="../pages/Users" class="nav-link">Yazarlar</a>
        </li>
      </ul>
      <div class="col-md-3 text-end">
        <div class="login-container">
          <button onclick="location.href= '../pages/Login-Register'" type="button" class="btn  me-2" style="background-color: lightseagreen; color: whitesmoke; font-size: 23px">Giriş</button>
          <button onclick="location.href= '../pages/Login-Register'" type="button" class="btn" style="background-color: lightseagreen; color: whitesmoke; font-size: 23px"> Kayıt</button>
        </div>
        <!--<div class="dropdown">
          <button class="btn dropdown-toggle" type="button" id="dropdownMenuButton1"
                  data-bs-toggle="dropdown" aria-expanded="false" style="background-color: lightseagreen; color: whitesmoke">
            Kullanıcı Adı
          </button>
          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
            <li><a class="dropdown-item" href="#">Bilgiler</a></li>
            <li><a class="dropdown-item" href="#">Yazılan Makaleler</a></li>
            <li><a class="dropdown-item" href="#">Çıkış Yap</a></li>
          </ul>
        </div>-->
      </div>
    </header>
    <!--<div class="row header-bottom text-center">
      <div class="col-6 d-flex">
        <div class="col">
          <p>Ana Sayfa</p>
        </div>
        <div class="col">
          <p>Makaleler</p>
        </div>
        <div class="col">
          <p>Yazarlar</p>
        </div>
        <div class="col">
          <p>Kategoriler</p>
        </div>
      </div>
      <div class="col-6 d-flex">
        <div class="col">
          <p>Search Bar</p>
        </div>
        <div class="col">
          <p>Giriş</p>
        </div>
        <div class="col">
          <p>Kayıt</p>
        </div>
      </div>
    </div>-->
  </div>
</template>

<script>

export default {
  name: "appHeader",
  data: function () {
    return {
      UserActive: false,
    }
  }
}
</script>

<style scoped>
.row {
  background-color: white;
}

a {
  color: #878687;
  font-size: 23px;
}

li > a:hover {
  text-decoration: underline 2px;
  text-decoration-color: seagreen;
  color: lightseagreen;
}
</style>