<template>
  <div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <a href="/home" class="navbar-brand">
        <img src="@/assets/logo.svg" alt="Logo" class="logo">
        <span class="site-name">Идентификация на КПП</span>
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

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400&display=swap');

body {
  margin: 0;
  font-family: 'Roboto', sans-serif;
}

.navbar {
  display: flex;
  align-items: center;
  gap: 50vh;
}

.navbar-container {
  display: flex;
  flex-direction: row;
}

.navbar-brand {
  display: flex;
  align-items: center;
}

.logo {
  width: 40px;
  height: 40px;
  margin-left: 20px;
  margin-right: 10px;
}

.site-name {
  font-size: 1.35rem;
  background: linear-gradient(90deg, #efeded, #0056b3, #007bff);
  background-size: 300%;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: gradient-animation 7s infinite;
  position: relative;
  font-family: 'Montserrat', sans-serif;
  margin-right: 20px;
}

.nav-link {
  font-size: 1rem;
  color: #ffffff;
  transition: color 0.3s, transform 0.3s;
  display: flex;
  align-items: center;
  font-family: 'Montserrat', sans-serif;
}

.nav-link .fa-icon {
  margin-right: 5px;
}

.nav-link:hover {
  color: #ff6b6b;
  transform: scale(1.05);
}

.nav-item {
  margin-left: 8vh;
}

.site-name::after {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background: rgba(80, 54, 54, 0.2);
  opacity: 0;
  transition: opacity 0.5s;
}

.site-name:hover::after {
  opacity: 1;
}

@keyframes gradient-animation {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}
</style>