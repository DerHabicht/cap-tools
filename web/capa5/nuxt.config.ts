// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  ssr: false,
  modules: [
    '@nuxt/image',
    '@nuxtjs/tailwindcss',
    '@primevue/nuxt-module',
    '@vueuse/nuxt',
    'nuxt-gravatar',
  ],
  primevue: {
    unstyled: true,
  }
})