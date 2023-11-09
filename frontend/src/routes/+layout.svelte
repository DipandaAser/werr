<script>
  import "carbon-components-svelte/css/all.css";
  import { Content, HeaderSearch, Theme } from "carbon-components-svelte";
  import {
    Header,
    HeaderAction,
    HeaderGlobalAction,
    HeaderActionLink,
    HeaderPanelDivider,
    HeaderPanelLink,
    HeaderPanelLinks,
    HeaderUtilities,
    SkipToContent,
  } from "carbon-components-svelte";
  import SettingsAdjust from "carbon-icons-svelte/lib/SettingsAdjust.svelte";
  import UserAvatarFilledAlt from "carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte";
  import Upload from "carbon-icons-svelte/lib/Upload.svelte";
  import {
    Blog,
    BrightnessContrast,
    Help,
    Login,
    LogoFacebook,
    LogoInstagram,
    LogoTwitter,
  } from "carbon-icons-svelte";

  let isSideNavOpen = false;
  let isSettingsOpen = false;
  let searchValue = "";
</script>

<Header company="Werr" platformName=".cm" bind:isSideNavOpen>
  <svelte:fragment slot="skip-to-content">
    <SkipToContent />
  </svelte:fragment>
  <HeaderUtilities>
    <!-- An action to upload -->
    <HeaderSearch bind:value={searchValue} />
    <HeaderActionLink icon={Upload} aria-label="Upload" />
    <HeaderAction aria-label="Settings" bind:isOpen={isSettingsOpen}>
      <HeaderPanelLinks>
        <HeaderPanelDivider>Account</HeaderPanelDivider>
        <HeaderPanelLink href=""><Login /> Login</HeaderPanelLink>
        <HeaderPanelDivider>Theme</HeaderPanelDivider>
        <HeaderPanelLink>
          <div class="theme-switcher">
            <BrightnessContrast />
            <span> Dark Theme</span>
            <Theme
              render="toggle"
              persist={true}
              toggle={{
                themes: ["g10", "g80"],
                labelA: "",
                labelB: "",
                hideLabel: true,
                size: "sm",
              }}
            />
          </div>
        </HeaderPanelLink>
        <HeaderPanelDivider>Resources</HeaderPanelDivider>
        <HeaderPanelLink href=""><Blog /> Read Blog</HeaderPanelLink>
        <HeaderPanelLink href=""><LogoInstagram /> Follow Us</HeaderPanelLink>
        <HeaderPanelLink href=""><Help /> Help</HeaderPanelLink>
        <HeaderPanelDivider>Follow Us</HeaderPanelDivider>
        <HeaderPanelLink href=""><LogoFacebook /> Facebook</HeaderPanelLink>
        <HeaderPanelLink href=""><LogoInstagram /> Instagram</HeaderPanelLink>
        <HeaderPanelLink href=""><LogoTwitter /> Twitter</HeaderPanelLink>
      </HeaderPanelLinks>
    </HeaderAction>
  </HeaderUtilities>
</Header>

<Content>
  <slot />
</Content>

<style>
  .theme-switcher {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .theme-switcher > span {
    margin-right: 40%;
  }
</style>
