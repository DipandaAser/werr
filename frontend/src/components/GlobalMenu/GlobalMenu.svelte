<script lang="ts">
  import { translatePath } from "$lib/i18n";
  import { CategoriesList } from "$lib/models/categorie";
  import Icon from "@iconify/svelte";
  import { authHandlers } from "../../stores/authStore";
  import { languageStore, updateLang } from "../../stores/languageStore";
  import { availableLanguageTags } from "$paraglide/runtime";
  import { page } from "$app/stores";
  export let onLangChange: () => void;
</script>

<div class="dropdownMenu">
  <div class="row">
    <div>
      <div class="dropdownMenuHeading">Media</div>
      {#each CategoriesList as categorie}
        <!-- style="margin-top: 0px;" -->
        <a
          class="dropdownMenuItem"
          href={translatePath(categorie.url, $languageStore.lang)}
        >
          <div class="icon">
            <Icon icon={categorie.icon} width="20" height="20" />
          </div>
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label>{categorie.label}</label>
        </a>
      {/each}
    </div>
    <div class="separator separator-vertical"></div>
    <div>
      <div class="dropdownMenuHeading">About</div>
      <a class="dropdownMenuItem" href="#">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>About Us</label>
      </a>
      <a class="dropdownMenuItem" href="#">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Terms of Services</label>
      </a>
      <a class="dropdownMenuItem" href="#">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Privacy Policy</label>
      </a>
      <a class="dropdownMenuItem" href="#">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Cookies Policy</label>
      </a>
    </div>
  </div>

  <div class="separator"></div>
  <div class="dropdownMenuHeading">
    <Icon icon="clarity:language-solid" width="20" height="20" />
    Language
  </div>
  {#each availableLanguageTags as lang}
    <a
      class="dropdownMenuItem"
      on:click={() => {
        onLangChange();
      }}
      href={translatePath($page.url.pathname, lang)}
      hreflang={lang}
    >
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <label>{lang}</label>
    </a>
  {/each}
</div>

<style lang="scss">
  .dropdownMenu {
    // min-width: 220px;
    display: flex;
    flex-direction: column;
    padding: 12px;
    margin-top: 0;
  }

  // get the first child of dropdownMenu who has the row class

  .dropdownMenuHeading {
    color: #fff;
    font-weight: 800;
    white-space: nowrap;
    padding: 8px 12px;
    border-radius: 8px;
    justify-content: start;
    display: flex;
    align-items: center;
  }

  .dropdownMenuItem {
    color: #ceced2;
    display: flex;
    align-items: center;
    padding: 8px 12px;
    text-decoration: none;
    flex-wrap: nowrap;
    white-space: nowrap;
    cursor: pointer;
    border-radius: 8px;
    -webkit-user-select: none;
    -ms-user-select: none;
    user-select: none;
    margin-top: 8px;
    margin-left: 4px;

    &:hover {
      background-color: #5b5b66;
      color: white;
    }
  }
  label {
    white-space: nowrap;
    cursor: pointer;
    font-size: 14px;
    font-weight: 400;
  }

  .separator {
    margin-top: 4px;
    margin-bottom: 2px;
    border: none;
    border-bottom: 1px solid hsla(0, 0%, 100%, 0.1);
  }

  .icon {
    margin-right: 4px;
  }

  @media (min-width: 769px) {
    .row {
      display: flex;
      flex-direction: row;
    }

    .separator-vertical {
      margin-left: 4px;
      margin-right: 4px;
      border: none;
      border-left: 1px solid hsla(0, 0%, 100%, 0.1);
    }
  }
</style>
