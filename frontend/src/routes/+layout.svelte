<script lang="ts">
  import "../app.scss";
  import "carbon-components-svelte/css/white.css";
  import { onMount } from "svelte";
  import { auth } from "$lib/firebase/firebase.client";
  import { authStore } from "../stores/authStore";
  import Header from "../components/Header/Header.svelte";
  import LoginSingUp from "../components/LoginSignUp/LoginSingUp.svelte";
  import { authPopupStore } from "../stores/authPopupStore";
  import { page } from "$app/stores";
  import {
    setLanguageTag,
    sourceLanguageTag,
    type AvailableLanguageTag,
  } from "$paraglide/runtime";
  import { browser } from "$app/environment";
  import { getTextDirection } from "$lib/i18n";
  import I18NHeader from "$lib/I18NHeader.svelte";
  import { updateLang } from "../stores/languageStore";
  onMount(() => {
    const unsubscribe = auth.onAuthStateChanged((user) => {
      console.log(user);
      authStore.update((/** @type {any} */ curr) => {
        return { ...curr, isLoading: false, currentUser: user };
      });
    });
    return unsubscribe;
  });
  //Determine the current language from the URL. Fall back to the source language if none is specified.
  $: lang = ($page.params.lang as AvailableLanguageTag) ?? sourceLanguageTag;

  $: updateLang(
    ($page.params.lang as AvailableLanguageTag) ?? sourceLanguageTag
  );

  //Set the language tag in the Paraglide runtime.
  //This determines which language the strings are translated to.
  //You should only do this in the template, to avoid concurrent requests interfering with each other.
  $: setLanguageTag(lang);

  //Determine the text direction of the current language
  $: textDirection = getTextDirection(lang);

  //Keep the <html> lang and dir attributes in sync with the current language
  $: if (browser) {
    document.documentElement.dir = textDirection;
    document.documentElement.lang = lang;
  }
</script>

<I18NHeader />

{#key lang}
  <Header />
  <slot />
  <!-- this is the login popup -->
  {#if $authPopupStore.isOpen}
    <LoginSingUp />
  {/if}
{/key}
