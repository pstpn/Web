<template>
  <div :class="['card card-container', cardValidationClass]">
    <div v-if="!created">
      <img
          id="profile-img"
          :src="imageUrl"
          sizes="(max-width:100px) 20px, 5vw"
          class="profile-img-card"
          alt="Not found"
      />
      <h5 class="card-title">Личная информация</h5>
      <Form class="card-body" ref="form" @submit="handleProfile || addField" :validation-schema="schema">
        <div class="form-group">
          <Field type="file" name="image" class="form-control-file" accept="image/jpeg" @change="handleImageUpload" />
          <ErrorMessage name="image" class="text-danger" />
        </div>
        <div class="form-group">
          <label for="documentType">Тип документа, удостоверяющего личность</label>
          <Field
              name="selectedType"
              as="select"
              class="form-control"
              :value="selectedType"
              @input="$emit('update:selectedType', $event.target.value)"
          >
            <option value="" disabled>Выберите тип документа</option>
            <option v-for="(documentType, index) in documentTypes" :key="index" :value="documentType">
              {{ documentType }}
            </option>
          </Field>
          <ErrorMessage name="selectedType" class="error-feedback" />
        </div>
        <div class="form-group">
          <label for="serialNumber">Серийный номер документа</label>
          <Field
              name="serialNumber"
              type="text"
              class="form-control"
              :value="serialNumber"
              @input="$emit('update:serialNumber', $event.target.value)"
          />
          <ErrorMessage name="serialNumber" class="error-feedback" />
        </div>
        <div class="form-group" v-for="(field, index) in documentFields" :key="index">
          <label :for="'documentField' + index">Поле документа {{ index + 1 }}</label>
          <div class="input-group">
            <input
                type="text"
                class="form-control"
                :value="field.type"
                @input="$emit('updateFieldType', { index, value: $event.target.value })"
                placeholder="Тип поля"
            >
            <input
                type="text"
                class="form-control"
                :value="field.value"
                @input="$emit('updateFieldValue', { index, value: $event.target.value })"
                placeholder="Значение поля"
            >
            <div class="input-group-append">
              <button @click="removeField(index)" class="btn btn-outline-secondary" type="button">Удалить</button>
            </div>
          </div>
        </div>
        <div class="form-group">
          <button @click="handleProfile" class="btn btn-primary col-md-12" :disabled="loading">
            <span v-show="loading" class="spinner-border spinner-border-sm"></span>
            Создать карточку
          </button>
          <button @click="addField" class="btn btn-dark col-md-12" :disabled="loading">
            <span v-show="loading" class="spinner-border spinner-border-sm"></span>
            Добавить поле
          </button>
        </div>
      </Form>
    </div>
    <div v-else>
      <img
          id="profile-img"
          :src="employeePhotoURL"
          sizes="(max-width:100px) 20px, 5vw"
          class="profile-img-card"
          alt="Not found"
      />
      <div class="card-body">
        <h5 class="card-title">Личная информация</h5>
        <p class="card-text">Тип документа: {{ profileInfo.documentType }}</p>
        <p class="card-text">Серийный номер документа: {{ profileInfo.serialNumber }}</p>
        <table class="table">
          <tbody>
          <tr v-for="(pair, index) in profileInfo.documentFields" :key="index">
            <td>{{ pair.type }}</td>
            <td>{{ pair.value }}</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div v-if="message" :class="created ? 'alert-success' : 'alert-danger'" class="alert">
      {{ message }}
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "ProfileCard",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  props: {
    schema: Object,
    created: Boolean,
    profileInfo: Object,
    employeePhotoURL: String,
    imageUrl: String,
    documentTypes: Array,
    selectedType: String,
    serialNumber: String,
    documentFields: Array,
    loading: Boolean,
    cardValidationClass: String,
    message: String,
  },
  methods: {
    addField() {
      this.$emit('addField');
    },
    removeField(index) {
      this.$emit('removeField', index);
    },
    handleImageUpload(event) {
      this.$emit('handleImageUpload', event);
    },
    handleProfile(profile) {
      this.$emit('handleProfile', profile);
    },
  },
};
</script>
