<script lang="ts">
  import { onMount } from "svelte";
  import Icon from "@iconify/svelte";
  import { openAuthPopup } from "../stores/authPopupStore";
  import { authStore } from "../stores/authStore";
  import { isCategoriesURLHome } from "$lib/models/categorie";
  const logoMobileDark = "/logo/mobile-black.svg";
  const logoMobileWhite = "/logo/mobile-white.svg";
  const logoDesktopDark = "/logo/desktop-black.svg";
  const logoDesktopWhite = "/logo/desktop-white.svg";
  let inScroll = false;

  onMount(() => {
    handleScrool();
    window.addEventListener("scroll", handleScrool);
  });

  function handleScrool() {
    const header = document.querySelector(".header") as HTMLElement;
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;

    const inMobile = window.innerWidth < 769;
    const registerBtn = document.querySelector(".register") as HTMLElement;
    const loginBtn = document.querySelector(".login") as HTMLElement;

    if (scrollTop > 0 || !isCategoriesURLHome(window.location.pathname)) {
      inScroll = true;
      header.style.backgroundColor = "white";
      header.style.color = "black";
      header.style.borderBottom = "1px solid #ebecf0";
      loginBtn.style.color = "#656f79";
      registerBtn.style.color = "black";
      if (inMobile) {
        const enterMobileSearch = document.querySelector(
          ".enter-mobile-search"
        ) as HTMLElement;
        enterMobileSearch.style.display = "flex";
      }
    } else {
      inScroll = false;
      header.style.backgroundColor = "transparent";
      header.style.color = "white";
      header.style.borderBottom = "";
      const enterMobileSearch = document.querySelector(
        ".enter-mobile-search"
      ) as HTMLElement;
      enterMobileSearch.style.display = "none";
      loginBtn.style.color = "white";
      registerBtn.style.color = "white";
    }
  }
</script>

<header class="header">
  <div class="sub-header-div">
    <div class="logo">
      <a href="/" aria-label="Werr">
        <picture>
          <source
            media="(min-width: 769px)"
            srcset={inScroll ? logoDesktopDark : logoDesktopWhite}
            type="image/svg+xml"
          />
          <img
            class="logo"
            src={inScroll ? logoMobileDark : logoMobileWhite}
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
    {#if !$authStore.currentUser}
      <div class="login-links">
        <button
          class="login"
          on:click={() => {
            openAuthPopup(false);
          }}>Log in</button
        >
        <button
          class="register"
          on:click={() => {
            openAuthPopup(true);
          }}>Join</button
        >
      </div>
    {/if}
    <div class="mobileMenu">
      <button type="button">
        <Icon
          icon="basil:menu-solid"
          color={inScroll ? "black" : "white"}
          width="28"
          height="28"
        />
      </button>
    </div>
  </div>
</header>

<style lang="scss">
  $hoverHeaderColor: rgba(182, 175, 175, 0.151);

  // .root {
  //   border: 1px solid red;
  // }

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
    padding: 0 16px;

    &.in-scrool {
      background-color: white;
      color: black;
      border-bottom: 1px solid #ebecf0;
    }

    .sub-header-div {
      max-width: 1830px;
      width: 100%;
      display: flex;
      align-items: center;
    }
  }

  .logo {
    flex: 0 0 auto;
  }

  .logo {
    width: 40px;
    height: 40px;
  }

  // large screen specific
  @media screen and (min-width: 769px) {
    .logo {
      width: 120px;
      height: 33px;
    }
    .enter-mobile-search {
      display: none;
    }
    .nav-search {
      display: block;
    }

    .header {
      padding: 0 32px;

      .sub-header-div {
        justify-content: space-between;
        padding: 0 32;
      }
    }
  }
</style>
