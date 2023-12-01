<script lang="ts">
  import Icon from "@iconify/svelte";
  export let icon: string = "";
  export let iconSize: number = 24;
  export let link: string = "/";
  export let isSelected: boolean = false;
  export let isSelectedColor: string;
  export let isNotSelectedColor: string;
  export let isSelectedIconColor: string;
  export let isNotSelectedIconColor: string;
  export let label: string = "";
  export let onSelect: () => void;
</script>

<div class="root">
  <a
    on:click={onSelect}
    href={link}
    style="background-color: {isSelected ? '#f4f4f4' : 'none'};
  "
  >
    {#if isSelected}
      <Icon
        {icon}
        class="icon"
        height={iconSize}
        width={iconSize}
        color={isSelected ? isSelectedIconColor : isNotSelectedIconColor}
      />
    {/if}
    <span
      class="label"
      style="color: {isSelected ? isSelectedColor : isNotSelectedColor};"
    >
      {label}
    </span>
  </a>
</div>

<style lang="scss">
  .root {
    cursor: pointer;
    padding-right: 12px;
  }

  a {
    border: none;
    border-radius: 24px;
    height: 40px;
    padding: 0 16px;
    display: inline-flex;
    align-items: center;
    justify-content: baseline;
    cursor: pointer;
    transition: all 0.1s ease-in;
    text-decoration: none;
    white-space: nowrap;
    font-weight: 600;

    // &:hover {
    //   background-color: #f4f4f4;
    // }
  }

  .icon {
    margin-right: 8px;
    font-size: 20px;
    margin-right: 8px;
  }

  .label {
    font-weight: 600;
    color: inherit;
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
