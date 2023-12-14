<script lang="ts">
  import Icon from "@iconify/svelte";
  import DropdownMenu from "./DropdownMenu.svelte";

  export let avatarUrl: string | undefined;
  export let name: string | undefined;
  export let closeFunction: () => void;
</script>

<div class="takeover">
  <span class="titleContainer">
    <div class="mobileDropdownHeader">
      <div class="container">
        <img alt="avatar" src={avatarUrl} class="avatar" />
      </div>
      {name}
    </div>
    <button class="closeButton" on:click={closeFunction}>
      <Icon icon="basil:cross-outline" color="white" width="24" height="24" />
    </button>
  </span>
  <hr class="divider" />
  <div class="contentContainer">
    <div>
      <DropdownMenu />
    </div>
  </div>
</div>

<style lang="scss">
  .takeover {
    top: 0;
    height: 100vh;
    height: calc(var(--vh, 1vh) * 100);
    border: 1px solid hsla(0, 0%, 100%, 0.1);
    background: #191b26;
    color: #fff;
    position: fixed;
    left: 0;
    width: 100vw;
    display: flex;
    flex-direction: column;
    z-index: 92;
  }

  .titleContainer {
    font-weight: 600;
    font-size: 16px;
    padding: 16px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;

    .closeButton {
      margin-left: auto;
      cursor: pointer;
      font-size: 20px;
      height: 24px;
      width: 24px;
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      border: none;
      background-color: unset;
      -webkit-appearance: button;
      appearance: button;
      overflow: visible;
      text-transform: none;
      margin: 0;
      font-weight: 400;

      &:hover {
        background-color: #838691;
      }
    }
  }

  .mobileDropdownHeader {
    display: flex;
    align-items: center;
    font-size: 16px;
    font-weight: 700;

    .container {
      width: 32px;
      height: 32px;
      margin-right: 8px;
      overflow: hidden;
      border-radius: 50%;
      display: inline-flex;

      .avatar {
        width: 100%;
        height: 100%;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        color: #656f79;
        -o-object-fit: cover;
        object-fit: cover;
        background: #fff;
      }
    }
  }

  .divider {
    margin: 0;
    border: none;
    border-bottom: 1px solid hsla(0, 0%, 100%, 0.1);
  }

  .contentContainer {
    height: calc(100% - 48px);
    overflow-y: auto;
    position: relative;

    div {
      color: #fff;
      background: #191b26;
      margin-right: 12px;
    }
  }
</style>
