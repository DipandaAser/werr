<script lang="ts">
  import "../app.scss";
  import "carbon-components-svelte/css/white.css";
  import { onMount } from "svelte";
  import { auth } from "$lib/firebase/firebase.client";
  import { authStore } from "../stores/authStore";
  import Header from "../components/Header/Header.svelte";
  import LoginSingUp from "../components/LoginSignUp/LoginSingUp.svelte";
  import { authPopupStore } from "../stores/authPopupStore";
  onMount(() => {
    const unsubscribe = auth.onAuthStateChanged((user) => {
      console.log(user);
      authStore.update((/** @type {any} */ curr) => {
        return { ...curr, isLoading: false, currentUser: user };
      });

      // if (
      //   browser &&
      //   !$authStore?.currentUser &&
      //   !$authStore.isLoading &&
      //   window.location.pathname !== routes.LOGIN
      // ) {
      //   window.location.href = routes.HOME;
      //   console.log($authStore.currentUser, $authStore.isLoading);
      // }
    });
    return unsubscribe;
  });
</script>

<Header></Header>

<slot />

<!-- this is the login popup -->
{#if $authPopupStore.isOpen}
  <LoginSingUp />
{/if}
