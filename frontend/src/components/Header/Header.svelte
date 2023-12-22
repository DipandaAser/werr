<script lang="ts">
  import { onMount } from "svelte";
  import Icon from "@iconify/svelte";
  import { openAuthPopup, toggleAuthPopup } from "../../stores/authPopupStore";
  import { authStore } from "../../stores/authStore";
  import { isCategoriesURLHome } from "$lib/models/categorie";
  import UserMenuButton from "./UserMenuButton.svelte";
  import AccountDropdownLarge from "../AccountDropdownMenu/AccountDropdownLarge.svelte";
  import AccountDropdownMobile from "../AccountDropdownMenu/AccountDropdownMobile.svelte";
  import * as multiLang from "$paraglide/messages";
  import { translatePath } from "$lib/i18n";
  import { routes } from "$lib";
  import { languageStore } from "../../stores/languageStore";
  import GlobalMenuMobile from "../GlobalMenu/GlobalMenuMobile.svelte";
  import GlobalMenuLarge from "../GlobalMenu/GlobalMenuLarge.svelte";
  import { page } from "$app/stores";
  import { afterNavigate } from "$app/navigation";
  const logoMobileDark = "/logo/mobile-black.svg";
  const logoMobileWhite = "/logo/mobile-white.svg";
  const logoDesktopDark = "/logo/desktop-black.svg";
  const logoDesktopWhite = "/logo/desktop-white.svg";
  let inScroll = false;
  let isUserMenuOpen = false;
  let isGlobalMenuOpen = false;
  let isHeaderOnTransparentMode =
    !inScroll && isCategoriesURLHome($page.url.pathname);

  onMount(() => {
    handleScrool(false);
    window.addEventListener("scroll", () => {
      handleScrool(false);
    });
    if (!isCategoriesURLHome(window.location.pathname)) {
      isHeaderOnTransparentMode = false;
      applyTransparentModeChanges();
    }
  });

  afterNavigate(() => {
    if (!isCategoriesURLHome(window.location.pathname)) {
      isHeaderOnTransparentMode = false;
      applyTransparentModeChanges();
    } else {
      handleScrool(true);
    }
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    inScroll = scrollTop > 0;
  });

  function applyTransparentModeChanges() {
    const header = document.querySelector(".header") as HTMLElement;
    const inMobile = window.innerWidth < 769;
    const registerBtn = document.querySelector(".register") as HTMLElement;
    const loginBtn = document.querySelector(".login") as HTMLElement;
    header.style.backgroundColor = "white";
    header.style.color = "black";
    header.style.borderBottom = "1px solid #ebecf0";
    if (!authStore.currentUser) {
      loginBtn.style.color = "#656f79";
      registerBtn.style.color = "black";
    }
    if (inMobile) {
      const enterMobileSearch = document.querySelector(
        ".enter-mobile-search"
      ) as HTMLElement;
      enterMobileSearch.style.display = "flex";
    }
  }

  function disableTransparentModeChanges() {
    const header = document.querySelector(".header") as HTMLElement;
    const inMobile = window.innerWidth < 769;
    const registerBtn = document.querySelector(".register") as HTMLElement;
    const loginBtn = document.querySelector(".login") as HTMLElement;
    header.style.backgroundColor = "transparent";
    header.style.color = "white";
    header.style.borderBottom = "";
    const enterMobileSearch = document.querySelector(
      ".enter-mobile-search"
    ) as HTMLElement;
    if (authStore.currentUser) {
      loginBtn.style.color = "white";
      registerBtn.style.color = "white";
    }
    enterMobileSearch.style.display = "none";
  }

  function handleScrool(refresh: boolean) {
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    inScroll = scrollTop > 0;
    console.log("url: ", window.location.pathname);
    if (!isCategoriesURLHome(window.location.pathname)) {
      isHeaderOnTransparentMode = false;
      console.log("not in home");
      return;
    }
    console.log("inscroll reached");

    if (inScroll || !isCategoriesURLHome(window.location.pathname)) {
      applyTransparentModeChanges();
    } else {
      disableTransparentModeChanges();
    }
  }

  function toggleUserMenu() {
    isUserMenuOpen = !isUserMenuOpen;
    if (isGlobalMenuOpen) {
      isGlobalMenuOpen = false;
    }
  }

  function toggleGlobalMenu() {
    isGlobalMenuOpen = !isGlobalMenuOpen;
    if (isUserMenuOpen) {
      isUserMenuOpen = false;
    }
  }

  function closeMenus() {
    isGlobalMenuOpen = false;
    isUserMenuOpen = false;
  }

  $: console.log("isHeaderOnTransparentMode: ", isHeaderOnTransparentMode);

  $: isHeaderOnTransparentMode =
    !inScroll && isCategoriesURLHome($page.url.pathname);
</script>

<header class="header {isHeaderOnTransparentMode ? '' : 'in-scrool'}">
  <div class="sub-header-div">
    <div class="logo">
      <a href={translatePath("/", $languageStore.lang)} aria-label="Werr">
        <picture>
          <source
            media="(min-width: 769px)"
            srcset={isHeaderOnTransparentMode
              ? logoDesktopWhite
              : logoDesktopDark}
            type="image/svg+xml"
          />
          <img
            class="logo"
            src={isHeaderOnTransparentMode ? logoMobileWhite : logoMobileDark}
            alt="Werr"
          />
        </picture>
      </a>
    </div>

    <div class="enter-mobile-search">
      <button
        type="button"
        on:click={() => {
          console.log("search clicked");
        }}
      >
        <Icon icon="basil:search-outline" color="black" />
      </button>
    </div>

    <div class="spacer" />
    <div class="explore">
      <button class="explore-btn" on:click={toggleGlobalMenu}>
        <span
          class="explore-label {isHeaderOnTransparentMode ? '' : 'in-scroll'}"
        >
          {multiLang.headerExplore()}
        </span>
        <Icon
          style="margin-left: 4px;"
          icon={isGlobalMenuOpen
            ? "basil:caret-up-solid"
            : "basil:caret-down-solid"}
          color={isGlobalMenuOpen
            ? "#00ab6b"
            : isHeaderOnTransparentMode
              ? "white"
              : "black"}
          width="28"
          height="28"
        />
      </button>
    </div>
    {#if isGlobalMenuOpen}
      <div class="global-menu-large">
        <GlobalMenuLarge onLangChange={toggleGlobalMenu} />
      </div>
    {/if}

    {#if !$authStore.currentUser}
      <div class="login-links">
        <button
          class="login"
          on:click={() => {
            closeMenus();
            openAuthPopup(false);
          }}>{multiLang.login()}</button
        >
        <button
          class="register"
          on:click={() => {
            closeMenus();
            openAuthPopup(true);
          }}>{multiLang.headerRegister()}</button
        >
      </div>
    {:else}
      <div class="userMenu">
        <UserMenuButton
          src={$authStore?.currentUser?.photoURL ?? undefined}
          onClick={toggleUserMenu}
        />
        {#if isUserMenuOpen && window.innerWidth >= 769}
          <AccountDropdownLarge
            name={$authStore?.currentUser?.displayName ?? undefined}
          />
        {/if}
      </div>
      {#if isUserMenuOpen && window.innerWidth < 769}
        <div>
          <AccountDropdownMobile
            avatarUrl={$authStore?.currentUser?.photoURL ?? undefined}
            name={$authStore?.currentUser?.displayName ?? undefined}
            closeFunction={toggleUserMenu}
          />
        </div>
      {/if}
    {/if}
    <div class="mobileMenu">
      <button type="button" on:click={toggleGlobalMenu}>
        <Icon
          icon="basil:menu-solid"
          color={inScroll ? "black" : "white"}
          width="28"
          height="28"
        />
      </button>
      {#if isGlobalMenuOpen}
        <div class="global-menu-mobile">
          <GlobalMenuMobile closeFunction={toggleGlobalMenu} />
        </div>
      {/if}
    </div>
    <a
      class="uploadButton"
      on:click={(e) => {
        // in the user is not connected yet, we open the auth popup instead of redirecting
        if (!$authStore.currentUser) {
          console.log("USER not connected yet. Please login");
          toggleAuthPopup();
          e.preventDefault();
        }
      }}
      href={translatePath(routes.ACCOUNT.MEDIA.UPLOAD, $languageStore.lang)}
    >
      <Icon icon="basil:cloud-upload-solid" width="20" height="20" />
      <span class="label">Upload</span>
    </a>
  </div>
</header>

<style lang="scss">
  $hoverHeaderColor: rgba(182, 175, 175, 0.151);

  // .root {
  //   border: 1px solid red;
  // }

  .headerOthers {
    border-bottom: 1px solid #ebecf0;
    background-color: white;
    color: black;
    border-bottom: 1px solid #ebecf0;
  }

  .mobileMenu {
    display: block;

    button {
      color: white;
      border: none;
      background-color: unset;
      cursor: pointer;
      width: 40px;
      height: 40px;
      border-radius: 20px;
      &:hover {
        background-color: $hoverHeaderColor;
      }
    }
  }

  .login-links {
    margin-left: auto;
    display: flex;
    align-items: center;

    button {
      padding: 0 16;
      min-width: 0;
      border-radius: 18px;
      height: 40px;
      display: inline-flex;
      align-items: center;
      margin: 5px;
      padding: 10px;
      cursor: pointer;
      transition: all 0.1s ease-in;
      text-decoration: none;
      white-space: nowrap;
      font-weight: 600;
      font-size: 14px;
      color: white;
    }

    .login {
      background-color: transparent;
      border: none;
      box-sizing: border-box;
      overflow: visible;
      text-transform: none;
      margin: 0;
      &:hover {
        background-color: $hoverHeaderColor;
      }
    }

    .register {
      background-color: rgba(255, 255, 255, 0.1);
      border: none;
      &:hover {
        border-color: white;
      }
    }
  }

  .enter-mobile-search {
    display: none;
    flex: 0 0 auto;
    margin-left: 8px;
    margin-right: 8px;

    button {
      color: white;
      border: none;
      background-color: unset;
      cursor: pointer;
      width: 40px;
      height: 40px;
      border-radius: 20px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 14px;
      font-weight: 400;
      overflow: visible;
      text-transform: none;
      margin: 0;
    }

    // add a hover effect to this button
    button:hover {
      background-color: $hoverHeaderColor;
    }
  }

  .nav-search {
    display: none;
  }

  .header {
    width: 100%;
    height: 64px;
    display: flex;
    justify-content: center;
    background-color: transparent;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 10;
    color: white;
    transition: background-color 0.2s;
    // padding: 0 16px;

    &.in-scrool {
      background-color: white;
      color: black;
      border-bottom: 1px solid #ebecf0;
    }

    .sub-header-div {
      padding: 0 16px;
      justify-content: flex-start;
      max-width: 1830px;
      width: 100%;
      display: flex;
      align-items: center;
    }
  }

  .logo {
    flex: 0 0 auto;
    width: 40px;
    height: 40px;
  }

  .userMenu {
    position: relative;
    display: flex;
    margin-left: 8px;
  }

  .spacer {
    flex: 1 1 auto;
  }

  .explore {
    display: none;
  }

  .global-menu-large {
    display: none;
  }

  .global-menu-mobile {
    display: flex;
  }

  .uploadButton {
    display: none;
  }

  // large screen specific
  @media screen and (min-width: 769px) {
    .logo {
      flex: 0 0 auto;
      width: 120px;
      height: 33px;
    }

    .uploadButton {
      margin-left: 12px;
      background: #00ab6b;
      color: #fff;
      border: none;
      border-radius: 24px;
      height: 40px;
      padding: 0 16px;
      box-sizing: border-box;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      transition: all 0.1s ease-in;
      text-decoration: none;
      white-space: nowrap;
      font-weight: 600;

      .label {
        font-weight: 600;
        color: inherit;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        margin-left: 3px;
      }
    }

    .global-menu-large {
      display: flex;
    }

    .global-menu-mobile {
      display: none;
    }

    .mobileMenu {
      display: none;
    }

    .enter-mobile-search {
      display: none;
    }
    .nav-search {
      display: block;
    }

    .header {
      padding: 0 16px;
      .sub-header-div {
        justify-content: space-between;
        padding: 0 0;
      }
    }

    .explore {
      display: flex;

      .explore-btn {
        padding: 0 16;
        min-width: 0;
        border-radius: 24px;
        height: 40px;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        transition: all 0.1s ease-in;
        margin: 5px;
        padding: 10px;
        cursor: pointer;
        transition: all 0.1s ease-in;
        text-decoration: none;
        white-space: nowrap;
        font-weight: 600;
        font-size: 14px;
        color: white;
        background-color: transparent;
        border: none;
        box-sizing: border-box;
        overflow: visible;
        text-transform: none;
        margin: 0;
        &:hover {
          background-color: $hoverHeaderColor;
        }

        .explore-label {
          font-weight: 600;
          color: inherit;
          display: flex;
          align-items: center;
          justify-content: center;

          &.in-scroll {
            color: black;
          }
        }
      }
    }
  }
</style>
