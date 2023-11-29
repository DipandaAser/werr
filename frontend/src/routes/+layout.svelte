<script lang="ts">
  import "../app.scss";
  import "carbon-components-svelte/css/white.css";
  import { Content } from "carbon-components-svelte";
  import { afterUpdate, onMount } from "svelte";
  import { auth } from "$lib/firebase/firebase.client";
  import { browser } from "$app/environment";
  import { authStore } from "../stores/authStore";
  import { routes } from "$lib";
  import Header from "../components/Header.svelte";
  import HomeWelcomeImage from "../components/HomeWelcome/HomeWelcomeImage.svelte";
  // onMount(() => {
  //   const unsubscribe = auth.onAuthStateChanged((user) => {
  //     console.log(user);
  //     authStore.update((/** @type {any} */ curr) => {
  //       return { ...curr, isLoading: false, currentUser: user };
  //     });

  //     if (
  //       browser &&
  //       !$authStore?.currentUser &&
  //       !$authStore.isLoading &&
  //       window.location.pathname !== routes.LOGIN
  //     ) {
  //       window.location.href = routes.HOME;
  //       console.log($authStore.currentUser, $authStore.isLoading);
  //     }
  //   });
  //   return unsubscribe;
  // });

  let isCurrentUrlHome = false;
</script>

<Header></Header>

<!-- <Content> -->
{#if isCurrentUrlHome}
  <HomeWelcomeImage></HomeWelcomeImage>
{/if}
<slot />
<!-- </Content> -->
