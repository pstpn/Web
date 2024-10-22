<template>
  <div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <a href="/home" class="navbar-brand">
        <router-link to="/home" class="nav-link">
          <img src="@/assets/logo.svg" alt="Logo" class="logo">
          <span class="site-name">Идентификация на КПП</span>
        </router-link>
      </a>
      <div class="navbar-container">
        <div class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link to="/home" class="nav-link">
              <font-awesome-icon icon="home" /> Домой
            </router-link>
          </li>
          <li v-if="showAdminBoard" class="nav-item">
            <router-link to="/find-employees" class="nav-link">
              <font-awesome-icon icon="user" /> Поиск сотрудников
            </router-link>
          </li>
        </div>
        <div v-if="!currentUser" class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link to="/register" class="nav-link">
              <font-awesome-icon icon="user-plus" /> Зарегистрироваться
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/login" class="nav-link">
              <font-awesome-icon icon="sign-in-alt" /> Войти
            </router-link>
          </li>
        </div>
        <div v-if="currentUser && !showAdminBoard" class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link to="/profile" class="nav-link">
              <font-awesome-icon icon="user" /> Профиль
            </router-link>
          </li>
        </div>
        <div v-if="currentUser" class="navbar-nav ml-auto">
          <li class="nav-item">
            <a class="nav-link" @click.prevent="logOut">
              <font-awesome-icon icon="sign-out-alt" /> Выйти
            </a>
          </li>
        </div>
      </div>
    </nav>
    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
import "@/styles/base.css";
import "@/styles/common.css";

export default {
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
    showAdminBoard() {
      return !!(this.currentUser && this.currentUser['isAdmin']);
    }
  },
  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  }
};
</script>