<script>
  import { routes } from "$lib";

  /** @type {import('./$types').PageData} */
  // export let data;
</script>

<div class="contentgg">
  <button
    on:click={() => {
      window.location.href = routes.LOGIN;
    }}
  >
    LOGIN
  </button>
</div>

<style>
  .contentgg {
    height: 2000px;
  }
</style>
