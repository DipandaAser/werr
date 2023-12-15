<script lang="ts">
  import { routes } from "$lib/index";
  import { CategoriesLabels, CategoriesList } from "$lib/models/categorie";
  import ChipsTab from "../../components/ChipTabs/ChipsTab.svelte";
  import HomeWelcomeImage from "../../components/HomeWelcome/HomeWelcomeImage.svelte";
  import type { Chip } from "../../components/ChipTabs/types";
  import {
    chipsStore,
    updateSelectedChip,
  } from "../../stores/homeCategoriesSelectionStore";
  import { updateSearchStore } from "../../stores/searchStore";
  import { translatePath } from "$lib/i18n";
  import { languageStore } from "../../stores/languageStore";

  /** @type {import('./$types').PageData} */
  // export let data;

  const welcomeTitle: string =
    "Stunning media library for your next video content";
  const description: string =
    "Over 4.2 thousand+ high quality images, videos and audio shared by our amazing community for your content creation journey.";

  const chips: Chip[] = CategoriesList.map((categorie) => {
    return {
      label: categorie.label,
      icon: categorie.icon,
      link: translatePath(categorie.url, $languageStore.lang),
    };
  });

  // insert home at position 0
  chips.unshift({
    label: "Home",
    icon: "basil:home-solid",
    link: translatePath(routes.HOME, $languageStore.lang),
  });

  function onSelectTab(tab: string): void {
    updateSelectedChip(tab);
  }

  updateSelectedChip("Home");
  updateSearchStore("", CategoriesLabels.IMAGES);
</script>

<svelte:head>
  <title>{description}</title>
</svelte:head>

<HomeWelcomeImage {welcomeTitle} {description} />

<ChipsTab {onSelectTab} selectedTab={$chipsStore.selected} {chips} />

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
