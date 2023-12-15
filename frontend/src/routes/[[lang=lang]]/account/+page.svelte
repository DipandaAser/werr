<script lang="ts">
  import { authHandlers } from "../../../stores/authStore";
  import { authStore } from "../../../stores/authStore";

  let email = "";

  authStore.subscribe((curr) => {
    console.log("CURR", curr);
    email = curr?.currentUser?.email as string;
  });
</script>

{#if $authStore.currentUser}
  <div>
    <h1>CURRENT USER: {email}</h1>
    <button on:click={authHandlers.logout}>Logout</button>
  </div>
{:else}
  <div>Loading....</div>
{/if}

<style>
  div {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  h1 {
    text-align: center;
  }
</style>
