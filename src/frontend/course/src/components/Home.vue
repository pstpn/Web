<template>
  <div class="gradient">
      <div class="home">
          <div class="header">
            <h2 class="header-header">Идентификация сотрудников компании на КПП</h2>
            <p class="header-text">Разрабатываемый программный продукт позволяет сотруднику СБ быстро и качественно организовывать мероприятия по обеспечению безопасности на объектах. В современном мире такая система является очень актуальной и востребованной, так как всегда необходимо поддерживать безопасность в абсолютно разных областях и местах. Поэтому организация безопасности объектов стала проще с нашим приложением.</p>
          </div>
      </div>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
  // Добавьте логику, если необходимо
};
</script>

<style scoped>
.home {
  display: flex;
  padding: 50px;
  backdrop-filter: blur(60px);
  height: 100%;
  width: 850px;
}

.header {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  max-width: 850px;
  margin-bottom: 80px;
  gap: 10px;
}

.header-header {
  font-size: 50px;
  background: linear-gradient(90deg, #efeded, #007bff);
  font-weight: 600;
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

.header-text {
  font-size: 18px;
  color: white;
}

.gradient{
  width: 100%;
  height: 100%;
  background-image: url("../assets/bg.jpeg");
  background-size: 100% 100%;
  position: absolute;
  left: 0;
  right: 0;
  z-index: -1;
}
</style>