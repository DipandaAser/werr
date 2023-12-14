<script lang="ts">
  import DropdownMenu from "./DropdownMenu.svelte";
  export let name: string | undefined;
</script>

<div class="dropdown">
  <div class="dropdownMenu">
    <div class="dropdownMenuHeading">{name}</div>
  </div>
  <DropdownMenu />
</div>

<style>
  .dropdown {
    color: #fff;
    background: #191b26;
    top: calc(100% + 4px);
    right: 0;
    position: absolute;
    min-width: 216px;
    border-radius: 8px;
    box-shadow:
      0 0 0 1px rgba(14, 19, 24, 0.07),
      0 2px 18px rgba(14, 19, 24, 0.25);
    z-index: 90;
  }

  .dropdownMenu {
    min-width: 220px;
    display: flex;
    flex-direction: column;
    padding: 12px;
    padding-bottom: 0;
    margin-bottom: 0;
  }

  .dropdownMenuHeading {
    color: #fff;
    font-weight: 800;
    white-space: nowrap;
    padding: 8px 12px;
    border-radius: 8px;
  }
</style>
