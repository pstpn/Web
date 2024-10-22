<template>
  <div class="card">
    <ProfileCard
        :schema="schema"
        :created="created"
        :profileInfo="profileInfo"
        :employeePhotoURL="employeePhotoURL"
        :imageUrl="imageUrl"
        :documentTypes="documentTypes"
        :selectedType="selectedType"
        :serialNumber="serialNumber"
        :documentFields="documentFields"
        :loading="loading"
        :cardValidationClass="cardValidationClass"
        :message="message"
        @addField="addField"
        @removeField="removeField"
        @update:selectedType="selectedType = $event"
        @update:serialNumber="serialNumber = $event"
        @updateFieldType="updateFieldType"
        @updateFieldValue="updateFieldValue"
        @handleImageUpload="handleImageUpload"
        @handleProfile="handleProfile"
    />
  </div>
</template>

<script>
import ProfileCard from "@/components/ProfileCard.vue";
import * as yup from "yup";

export default {
  name: "Employee",
  components: {
    ProfileCard,
  },
  data() {
    const schema = yup.object().shape({
      image: yup.mixed().required("Загрузите ваше фото или изображение документа, содержащего фото"),
      serialNumber: yup.string().required("Введите серийный номер документа!"),
    });
    return {
      loading: false,
      message: "",
      schema,
      documentTypes: [ "Паспорт", "Водительские права" ],
      documentFields: [],
      selectedType: "",
      serialNumber: "",
      imageUrl: "//ssl.gstatic.com/accounts/ui/avatar_2x.png",
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn
    },
    created() {
      if (!this.profileInfo) {
        this.$store.dispatch("employee/getProfile").then(
            () => {
              this.$store.dispatch("employee/getEmployeePhoto");
            },
            (error) => {
              if (error.response && error.response.status === 404) {
                this.$store.state.employee.profile = null;
              }
            }
        )
      }
      return this.profileInfo;
    },
    profileInfo() {
      return this.$store.state.employee.profile;
    },
    employeePhotoURL() {
      return this.$store.state.employee.photoURL;
    },
    cardValidationClass() {
      if (!this.created)
        return 'not-created';
      if (this.profileInfo == null)
        return 'card-invalid';
      return this.profileInfo.isConfirmed ? 'card-valid' : 'card-invalid';
    }
  },
  methods: {
    addField() {
      this.documentFields.push({ type: "", value: "" });
    },
    removeField(index) {
      this.documentFields.splice(index, 1);
    },
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.image = file;

        const reader = new FileReader();
        reader.onload = (e) => {
          this.imageUrl = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    handleProfile(profile) {
      if (!this.created) {
        this.$refs.form.validate().then(success => {
          if (success) {
            this.fillProfile(profile);
          }
        });
      }
    },
    fillProfile(profile) {
      let user = JSON.parse(localStorage.getItem('user'));

      const formData = new FormData();
      formData.append('image', this.image);
      formData.append('serialNumber', this.$refs.form.values.serialNumber);
      formData.append('documentType', this.selectedType);

      let fields = '';
      for (let i = 0; i < this.documentFields.length; ++i)
        fields += this.documentFields[i].type + ',' + this.documentFields[i].value + ';';
      formData.append('documentFields', fields.slice(0, -1));

      this.loading = true;

      this.$store.dispatch("employee/fillProfile", formData).then(
          () => {
            this.loading = false;
            window.location.reload();
          },
          (error) => {
            this.loading = false;
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.fillProfile(profile);
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
  },
};
</script>