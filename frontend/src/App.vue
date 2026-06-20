<script lang="ts" setup>
import MapEditor from './components/MapEditor.vue';
</script>

<template>
  <div class="app">
    <h1 class="title">Room Guide</h1>
    <MapEditor></MapEditor>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body, #app {
  width: 100%;
  height: 100%;
}

.app {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}

.title {
  padding: 8px 16px;
  flex-shrink: 0;
}
</style>