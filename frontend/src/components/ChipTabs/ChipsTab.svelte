<script lang="ts">
  import Chips from "./Chips.svelte";
  import type { Chip } from "./types";

  const defaultIconSelectedColor: string = "#00ab6b";
  const defaultLabelSelectedColor: string = "black";
  const defaultLabelNotSelectedColor: string = "#5B5B66;";
  const defaultIconNotSelectedColor: string = "#5B5B66;";

  export let chips: Chip[] = [];
  export let isSelectedColor: string = defaultLabelSelectedColor;
  export let isNotSelectedColor: string = defaultLabelNotSelectedColor;
  export let isSelectedIconColor: string = defaultIconSelectedColor;
  export let isNotSelectedIconColor: string = defaultIconNotSelectedColor;
  export let iconSize: number = 24;
  export let selectedTab: string = chips[0].label;
  function onSelect(tab: string): void {
    selectedTab = tab;
  }
  export let onSelectTab: (tab: string) => void;

  onSelectTab(selectedTab);
</script>

<div class="container">
  <div class="tabs">
    {#each chips as chip}
      <Chips
        icon={chip.icon}
        isSelected={selectedTab === chip.label}
        onSelect={() => {
          onSelect(chip.label);
          onSelectTab(chip.label);
        }}
        {iconSize}
        {isSelectedColor}
        {isNotSelectedColor}
        {isSelectedIconColor}
        {isNotSelectedIconColor}
        label={chip.label}
        link={chip.link}
      />
    {/each}
  </div>
</div>

<style lang="scss">
  .container {
    padding: 24px 16px;
    margin-left: -16px;
    // margin-right: -16px;
    border-bottom: 1px solid #ebecf0;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .tabs {
    margin-left: -8px;
    padding-left: 16px;
    margin-right: -16px;
    padding-right: 16px;
    display: flex;
    overflow-x: scroll;
    -ms-overflow-style: none;
    scrollbar-width: none;
  }

  div::-webkit-scrollbar {
    width: 0; /* Vertical scrollbar width */
    height: 0; /* Horizontal scrollbar height */
  }
</style>
