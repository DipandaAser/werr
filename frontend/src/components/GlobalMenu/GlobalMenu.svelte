<script>
  import { authHandlers } from "../../stores/authStore";
</script>

<div class="dropdownMenu">
  <div class="dropdownMenuHeading">Media</div>
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
    <label>Statistiques</label>
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
  <div class="dropdownMenuHeading">About</div>
</div>

<style lang="scss">
  .dropdownMenu {
    // min-width: 220px;
    display: flex;
    flex-direction: column;
    padding: 12px;
    margin-top: 0;
  }

  .dropdownMenuHeading {
    color: #fff;
    font-weight: 800;
    white-space: nowrap;
    padding: 8px 12px;
    border-radius: 8px;
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
