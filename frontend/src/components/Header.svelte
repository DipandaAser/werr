<script lang="ts">
  import { onMount } from "svelte";
  import Icon from "@iconify/svelte";
  const logoMobileClass = "logo-mobile";
  const inScrollClass = "in-scrool";
  let inScroll = false;
  const outScrollClass = "out-scrool";

  onMount(() => {
    window.addEventListener("scroll", () => {
      const header = document.querySelector(".header") as HTMLElement;
      const scrollTop =
        window.pageYOffset || document.documentElement.scrollTop;

      const inMobile = window.innerWidth < 769;
      const registerBtn = document.querySelector(".register") as HTMLElement;
      const loginBtn = document.querySelector(".login") as HTMLElement;

      if (scrollTop > 0) {
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
    });
  });
</script>

<header class="header root">
  <div class="sub-header-div">
    <div class="logo">
      <a href="/" aria-label="Werr">
        <svg
          class={logoMobileClass}
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          width="40"
          height="40"
          viewBox="0 0 40 40"
        >
          <defs>
            <clipPath id="clip-LOGO_40">
              <rect width="40" height="40" />
            </clipPath>
          </defs>
          <g id="LOGO_40" data-name="LOGO 40" clip-path="url(#clip-LOGO_40)">
            <rect
              id="Rectangle_4"
              data-name="Rectangle 4"
              width="40"
              height="40"
              rx="9"
              fill="#fff"
            />
            <text
              id="w"
              transform="translate(12 16)"
              font-size="16"
              font-family="Delius-Regular, Delius"
              ><tspan x="-6.288" y="0">w</tspan></text
            >
            <text
              id="e"
              transform="translate(30 16)"
              font-size="16"
              font-family="Delius-Regular, Delius"
              ><tspan x="-4.312" y="0">e</tspan></text
            >
            <text
              id="r"
              transform="translate(12 36)"
              font-size="16"
              font-family="Delius-Regular, Delius"
              ><tspan x="-3.264" y="0">r</tspan></text
            >
            <text
              id="r-2"
              data-name="r"
              transform="translate(30 36)"
              font-size="16"
              font-family="Delius-Regular, Delius"
              ><tspan x="-3.264" y="0">r</tspan></text
            >
          </g>
        </svg>
        <!-- TODO: Add logo desktop -->
      </a>
    </div>
    <div class="enter-mobile-search">
      <button
        type="button"
        on:click={() => {
          console.log("search clicked");
        }}
      >
        <span> <Icon icon="basil:search-outline" color="black" /> </span>
      </button>
    </div>
    <div class="login-links">
      <button class="login">Log in</button>
      <button class="register">Join</button>
    </div>
    <div class="mobileMenu">
      <button type="button">
        <span>
          <Icon icon="basil:menu-solid" color={inScroll ? "black" : "white"} />
        </span>
      </button>
    </div>
  </div>
</header>

<style lang="scss">
  $hoverHeaderColor: rgba(182, 175, 175, 0.151);

  .root {
    border: 1px solid red;
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

  .logo-mobile {
    width: 40px;
    height: 40px;
  }

  svg {
    border: 1px solid red;
  }

  // large screen specific
  @media screen and (min-width: 769px) {
    .logo-mobile {
      display: none;
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
