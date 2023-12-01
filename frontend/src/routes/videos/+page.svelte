<script lang="ts">
  import { routes } from "$lib/index";
  import { CategoriesLabels, CategoriesList } from "$lib/models/categorie";
  import ChipsTab from "../../components/ChipTabs/ChipsTab.svelte";
  import type { Chip } from "../../components/ChipTabs/types";
  import HomeWelcomeImage from "../../components/HomeWelcome/HomeWelcomeImage.svelte";
  import {
    chipsStore,
    updateSelectedChip,
  } from "../../stores/homeCategoriesSelectionStore";
  import { updateSearchStore } from "../../stores/searchStore";

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
      link: categorie.url,
    };
  });

  // insert home at position 0
  chips.unshift({
    label: "Home",
    icon: "basil:home-solid",
    link: routes.HOME,
  });

  function onSelectTab(tab: string): void {
    updateSelectedChip(tab);
  }

  updateSelectedChip(CategoriesLabels.VIDEOS);
  updateSearchStore("", $chipsStore.selected);
</script>

<svelte:head>
  <title>{$chipsStore.selected}</title>
</svelte:head>

<HomeWelcomeImage {welcomeTitle} {description} />

<ChipsTab {onSelectTab} selectedTab={$chipsStore.selected} {chips} />

<div class="contentgg">{$chipsStore.selected}</div>

<style>
  .contentgg {
    height: 2000px;
    text-align: center;
  }
</style>
