<template>
  <div class="search-container">
    <SearchBar
        :searchQuery="searchQuery"
        :searchBy="searchBy"
        :sortDirection="sortDirection"
        @update:searchQuery="searchQuery = $event"
        @update:searchBy="searchBy = $event"
        @update:sortDirection="sortDirection = $event"
        @search="searchEmployees"
    />
    <div v-if="showSearchResults" class="search-results">
      <div v-for="infoCard in searchResults" :key="infoCard.id" @click="infoCard.post !== 'Сотрудник СБ' ? viewEmployeeCard(infoCard) : mock()" class="search-item">
        <div class="employee-info">
          <div class="employee-details">
            <div class="employee-fullName">{{ infoCard.fullName }}</div>
            <div class="employee-phoneNumber">{{ infoCard.phoneNumber }}</div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showEmployeeCard" :class="['employee-card']">
      <div :class="['card card-container', cardValidationClass]">
        <img
            id="profile-img"
            :src="employeePhotoURL"
            sizes="(max-width:50px) 10px, 5vw"
            class="profile-img-card"
            alt="Not found"
        />
        <div class="card-body">
          <h5 class="card-title">{{ selectedEmployee.fullName }}</h5>
          <p class="card-subtitle" style="font-weight: bold;">Основная информация</p>
          <p class="card-text">Номер телефона: {{ selectedEmployee.phoneNumber }}</p>
          <p class="card-text">Должность: {{ selectedEmployee.post }}</p>
          <p v-if="selectedEmployeeDocument === null" class="card-subtitle" style="color: red; font-weight: bold;">Документ, удостоверяющий личность не найден</p>
          <p v-else class="card-subtitle" style="font-weight: bold;">Документ, удостоверяющий личность</p>
          <p v-if="selectedEmployeeDocument != null" class="card-text">Тип документа: {{ selectedEmployeeDocument.data.documentType }}</p>
          <p v-if="selectedEmployeeDocument != null" class="card-text">Серийный номер документа: {{ selectedEmployeeDocument.data.serialNumber }}</p>
          <p v-if="selectedEmployeeDocument != null" class="card-subtitle" style="font-weight: bold;">Поля документа</p>
          <table v-if="selectedEmployeeDocument != null" class="table">
            <tbody>
            <tr v-for="(pair, index) in selectedEmployeeDocument.fields" :key="index">
              <td>{{ pair.type }}</td>
              <td>{{ pair.value }}</td>
            </tr>
            </tbody>
          </table>
          <button v-if="!this.selectedEmployee.isConfirmed" @click="confirmInfoCard" class="btn btn-dark col-md-12">Подтвердить данные карточки</button>
          <div v-if="this.selectedEmployee.isConfirmed">
            <p class="card-subtitle" style="font-weight: bold;">Управление проходами</p>
            <button @click="addPassage('Вход')" class="btn btn-primary btn-dark col-md-6">Зафиксировать вход</button>
            <button @click="addPassage('Выход')" class="btn btn-primary btn-dark col-md-6">Зафиксировать выход</button>
            <table v-if="selectedEmployeeDocument != null" class="table">
              <tbody>
              <tr v-for="(passage, index) in passages[this.selectedEmployee.id]" :key="index">
                <td>{{ passage.type }}</td>
                <td>{{ passage.time }}</td>
              </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import "@/styles/admin.css";
import "@/styles/common.css";
import SearchBar from "@/components/SearchBar.vue";

export default {
  name: "AdminSearch",
  components: {
    SearchBar,
  },
  data() {
    return {
      searchQuery: '',
      searchBy: 'full_name',
      sortDirection: 'ASC',
      searchResults: [],
      showSearchResults: false,
      selectedEmployee: null,
      selectedEmployeeDocument: null,
      showEmployeeCard: false,
      passages: {},
    };
  },
  computed: {
    employeePhotoURL() {
      return (this.selectedEmployee &&  this.selectedEmployee.photoURL)
          ? this.selectedEmployee.photoURL
          : "//ssl.gstatic.com/accounts/ui/avatar_2x.png";
    },
    cardValidationClass() {
      return this.selectedEmployee.isConfirmed ? 'card-valid' : 'card-invalid';
    }
  },
  methods: {
    searchEmployees({ searchQuery, searchBy, sortDirection }) {
      this.searchQuery = searchQuery;
      this.searchBy = searchBy;
      this.sortDirection = sortDirection;

      if (this.searchQuery.length > 0) {
        this.showSearchResults = true;
        const { searchQuery, searchBy, sortDirection } = this;
        this.$store.dispatch('employee/getEmployees', { searchQuery, searchBy, sortDirection }).then(
            (employees) => {
              this.searchResults = employees;
            },
            (error) => {
              if (error.response && error.response.status === 404) {
                this.$store.state.employee.profile = null;
              }
              if (error.response && error.response.status === 401) {
                this.$store.dispatch('auth/refreshTokens', user).then(
                    response => {
                      this.searchEmployees();
                    },
                    (error) => {
                      if (error.response && error.response.status === 401) {
                        this.$store.dispatch('auth/logout');
                        this.$router.push('/login');
                      } else {
                        this.message = error.message + ": " + error.response.data.error;
                      }
                    }
                );
              } else {
                this.message = error.message + ": " + error.response.data.error;
              }
            }
        );
      } else {
        this.showEmployeeCard = false;
        this.showSearchResults = false;
        this.searchResults = [];
      }
    },
    viewEmployeeCard(infoCard) {
      let user = JSON.parse(localStorage.getItem('user'));

      const id = infoCard.id;
      this.selectedEmployee = infoCard;
      this.showEmployeeCard = true;
      this.$store.dispatch('employee/getEmployee', id).then(
          (employee) => {
            this.selectedEmployeeDocument = employee.document;
            if (employee.passages != null) {
              this.passages[id] = employee.passages;
            }
            this.$store.dispatch('employee/getEmployeeInfoCardPhoto', id).then(
                (photoURL) => {
                  this.selectedEmployee.photoURL = photoURL;
                }
            );
          },
          (error) => {
            if (error.response && error.response.status === 404) {
              this.selectedEmployeeDocument = null;
            } else if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.viewEmployeeCard(infoCard);
                  },
                  (error) => {
                    if (error.response && error.response.status === 401) {
                      this.$store.dispatch('auth/logout');
                      this.$router.push('/login');
                    } else {
                      this.message = error.message + ": " + error.response.data.error;
                    }
                  }
              );
            } else {
              this.message = error.message + ": " + error.response.data.error;
            }
          }
      );
    },
    addPassage(type) {
      const now = new Date();
      const formattedTime = `${this.padZero(now.getHours())}:${this.padZero(now.getMinutes())}:${this.padZero(now.getSeconds())} (${this.padZero(now.getDate())}.${this.padZero(now.getMonth() + 1)}.${now.getFullYear()})`;
      this.$store.dispatch('employee/createEmployeePassage', {
        infoCardID: this.selectedEmployee.id,
        type: type,
        time: now,
      }).then(() => {
        this.addPassageToDictionary(type, formattedTime);
      });
    },
    addPassageToDictionary(type, time) {
      const id = this.selectedEmployee.id;
      if (!this.passages.hasOwnProperty(id)) {
        this.passages[id] = [];
      }
      this.passages[id].push({ type: type, time: time });
    },
    padZero(number) {
      return number < 10 ? '0' + number : number;
    },
    confirmInfoCard() {
      let user = JSON.parse(localStorage.getItem('user'));

      this.$store.dispatch('employee/confirmEmployeeCard', this.selectedEmployee.id).then(
          (response) => {
            this.selectedEmployee.isConfirmed = true;
          },
          (error) => {
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.confirmInfoCard();
                  },
                  (error) => {
                    if (error.response && error.response.status === 401) {
                      this.$store.dispatch('auth/logout');
                      this.$router.push('/login');
                    } else {
                      this.message = error.message + ": " + error.response.data.error;
                    }
                  }
              );
            } else {
              this.message = error.message + ": " + error.response.data.error;
            }
          }
      );
    },
    mock() {

    }
  }
};
</script>
