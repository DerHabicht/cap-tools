<script setup lang="ts">
const route = useRoute()

await navigateTo({
  path: `/units/${route.params.unit}/ood`,
})
</script>
