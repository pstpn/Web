<template>
  <div class="search-bar">
    <div class="search-input-container">
      <input
          type="text"
          v-model="localSearchQuery"
          @input="emitSearch"
          placeholder="Поиск сотрудников"
          class="search-input"
      />
      <div class="search-input-icon">
        <i class="fas fa-search"></i>
      </div>
    </div>
    <div class="filter-dropdown" @mouseover="isDropdownVisible = true" @mouseleave="isDropdownVisible = false">
      <div class="dropdown-trigger">
        <font-awesome-icon icon="fa-filter"/>
      </div>
      <transition name="fade">
        <div v-if="isDropdownVisible" class="dropdown-content">
          <select v-model="localSearchBy" @change="emitSearch" class="filter-select">
            <option value="full_name">ФИО</option>
            <option value="phone_number">Номер телефона</option>
          </select>
        </div>
      </transition>
    </div>
    <div class="filter-dropdown" @mouseover="isSortDropdownVisible = true" @mouseleave="isSortDropdownVisible = false">
      <div class="dropdown-trigger">
        <font-awesome-icon icon="fa-sort"/>
      </div>
      <transition name="fade">
        <div v-if="isSortDropdownVisible" class="dropdown-content">
          <select v-model="localSortDirection" @change="emitSearch" class="filter-select">
            <option value="ASC">По возрастанию</option>
            <option value="DESC">По убыванию</option>
          </select>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import "@/styles/common.css";
import "@/styles/admin.css";

export default {
  props: {
    searchQuery: {
      type: String,
      default: ''
    },
    searchBy: {
      type: String,
      default: 'full_name'
    },
    sortDirection: {
      type: String,
      default: 'ASC'
    }
  },
  data() {
    return {
      localSearchQuery: this.searchQuery,
      localSearchBy: this.searchBy,
      localSortDirection: this.sortDirection,
      isDropdownVisible: false,
      isSortDropdownVisible: false
    };
  },
  watch: {
    // Синхронизируем локальные данные с prop-ами, если их изменят снаружи
    searchQuery(newVal) {
      this.localSearchQuery = newVal;
    },
    searchBy(newVal) {
      this.localSearchBy = newVal;
    },
    sortDirection(newVal) {
      this.localSortDirection = newVal;
    }
  },
  methods: {
    emitSearch() {
      this.$emit('update:searchQuery', this.localSearchQuery);
      this.$emit('update:searchBy', this.localSearchBy);
      this.$emit('update:sortDirection', this.localSortDirection);
      this.$emit('search', {
        searchQuery: this.localSearchQuery,
        searchBy: this.localSearchBy,
        sortDirection: this.localSortDirection
      });
    }
  }
};
</script>
