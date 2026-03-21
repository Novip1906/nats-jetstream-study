<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
</script>

<style>
body { font-family: Arial, sans-serif; background: #f0f2f5; display: flex; justify-content: center; padding: 50px; }
.container { background: white; padding: 20px; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); width: 100%; max-width: 500px; margin: 0 auto;}
input { width: calc(100% - 22px); padding: 10px; margin-bottom: 10px; border: 1px solid #ccc; border-radius: 4px; }
.buttons { display: flex; gap: 10px; margin-bottom: 20px; flex-wrap: wrap; }
button { padding: 10px 15px; border: none; border-radius: 4px; cursor: pointer; color: white; width: 100%; }
.btn-send { background: #007bff; }
.btn-send:hover { background: #0056b3; }
.btn-receive { background: #28a745; }
.btn-receive:hover { background: #218838; }
.message-list { list-style: none; padding: 0; }
.message-item { background: #e9ecef; margin-bottom: 5px; padding: 10px; border-radius: 4px; display: flex; justify-content: space-between; align-items: center; }
.btn-delete { background: #dc3545; padding: 5px 10px; font-size: 12px; width: auto; }
.btn-delete:hover { background: #c82333; }
.btn-send-fast { background: #17a2b8; }
.btn-send-fast:hover { background: #138496; }
.error-text { color: red; margin-bottom: 10px; }
.nav-link { color: #007bff; cursor: pointer; text-decoration: underline; margin-top: 10px; display: inline-block; }
</style>
