<script>
  import { authHandlers } from "../../stores/authStore";
</script>

<div class="dropdownMenu">
  <a class="dropdownMenuItem" style="margin-top: 0px;" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Mon profil</label>
  </a>
  <a class="dropdownMenuItem" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Mes médias</label>
  </a><a class="dropdownMenuItem" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Upload</label>
  </a>
  <a class="dropdownMenuItem" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Collections</label>
  </a>
  <a class="dropdownMenuItem" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Abonnements</label>
  </a>
  <a class="dropdownMenuItem" style="margin-bottom: 4px;" href="#">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Paramètres</label>
  </a>
  <div class="separator"></div>
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div on:click={authHandlers.logout} class="dropdownMenuItem">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Déconnexion</label>
  </div>
</div>

<style lang="scss">
  .dropdownMenu {
    // min-width: 220px;
    display: flex;
    flex-direction: column;
    padding: 12px;
    margin-top: 0;
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
    margin: 0;
    border: none;
    border-bottom: 1px solid hsla(0, 0%, 100%, 0.1);
  }
</style>
