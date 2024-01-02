<script lang="ts">
  import { CategoriesIcons, CategoriesList } from "$lib/models/categorie";
  import "../../../../../app.scss";
  import Icon from "@iconify/svelte";
  import * as multiLang from "$paraglide/messages";
  import { onMount } from "svelte";
  import EditInfo from "../../../../../components/Media/Upload/EditInfo.svelte";
  import { NewApiClient } from "$lib/api/apiClient";
  import { MediaAPI } from "$lib/api/media";
  import { invalidateAll } from "$app/navigation";
  import { authStore } from "../../../../../stores/authStore";
  import { uploadFailedMessage } from "$paraglide/messages/en";

  /** @type {import('./$types').PageData} */
  export let data;

  let FileToUpload: File | null = null;
  let fileName: String = "";
  let tags: String[] = [];
  let description: String = "";

  let showNotif = false;
  let notifMessage = "";
  let isNotifError = false;

  onMount(() => {
    setupFileSelectListerner();
  });

  function setupFileSelectListerner() {
    const input = document.querySelector(
      'input[type="file"]'
    ) as HTMLInputElement;
    input.onchange = () => {
      if (input.files && input.files.length > 0) {
        FileToUpload = input.files[0];
        fileName = FileToUpload?.name ?? "";
      }
    };
  }

  function upload() {
    const input = document.querySelector(
      'input[type="file"]'
    ) as HTMLInputElement;
    input.click();
    console.log("upload");
  }

  function submit() {
    console.log("Submission Started");
    $authStore.currentUser
      ?.getIdToken()
      .then((token) => {
        console.log("Token is: ", token);
        NewApiClient(token)
          .Media.ReqUpload(fileName, FileToUpload)
          .then((res) => {
            showNotif = true;
            isNotifError = false;
            console.log("Upload done. ID: ", res.id);
          })
          .catch((err) => {
            showNotif = true;
            isNotifError = true;
            console.log("Error is: ", err);
          });
      })
      .catch((err) => {
        showNotif = true;
        isNotifError = true;
        console.log("Error is: ", err);
      });
  }

  function cancel() {
    window.location.href = window.location.href;
    // FileToUpload = null;
    // invalidateAll();
  }

  $: console.log(FileToUpload);
</script>

<div class="page">
  <div class="container-1">
    {#if !FileToUpload}
      <div class="zeroUploadsContainer">
        <h1>Upload your media</h1>
        <h2>Join our community of contributors by uploading media!</h2>
        <h2>
          Learn more about the Pixabay <a
            href="/service/license-summary/"
            target="_blank">Content license.</a
          >
        </h2>
        <div role="presentation" class="container-2">
          <div class="container-3">
            <div class="mediaIconsContainer">
              <div class="videoIconContainer">
                <Icon
                  class="icon video"
                  icon={CategoriesIcons.VIDEOS}
                  color="#909cf2"
                  height="32"
                  width="32"
                />
              </div>
              <div class="imageIconContainer">
                <Icon
                  class="icon image"
                  icon={CategoriesIcons.IMAGES}
                  color="#e3ab6c"
                  height="32"
                  width="32"
                />
              </div>
              <div class="musicIconContainer">
                <Icon
                  class="icon audio"
                  icon={CategoriesIcons.AUDIOS}
                  color="#67a871;"
                  height="32"
                  width="32"
                />
              </div>
            </div>
          </div>
          <input
            accept="audio/*,.mp3,.mpga,.mp2,.mp2a,.m2a,.m3a,.wav,.aac,.flac,.aif,.aiff,.aifc,.m4a,.mp4a,.m4b,image/*,.gif,.jpg,.jpeg,.jpe,.jif,.jfif,.jfi,.png,.svg,.svgz,.ai,.psd,video/*,.mp4,.mp4v,.mpg4,.m4v,.mpeg,.mpg,.mpe,.m1v,.m2v,.mov,.qt,.avi"
            type="file"
            style="visibility: hidden; position: absolute;"
            tabindex="-1"
          /><span class="promptText"
            ><button on:click={upload} class="buttonBase primaryButton base">
              <span class="label">{multiLang.browseFiles()}</span></button
            ></span
          >
          <div class="infoSection">
            <!-- TODO: get the number of upload remaining in the user acount store -->
            <span class="uploadsRemainingText"
              >7 {multiLang.uploadPage_UploadRemain()}</span
            >
            <div class="tooltipTriggerWrapper">
              <Icon icon="octicon:question-24"></Icon>
              <span class="tooltiptext">
                {multiLang.uploadPage_UploadRemainDetails()}
              </span>
            </div>
          </div>
        </div>
        <div class="guidelinesContainer">
          {#each CategoriesList as categorie}
            <div>
              <Icon
                class="icon"
                style="margin-right: 16px;"
                color="#656f79"
                icon={categorie.icon}
                height="36"
                width="36"
              />
              <span class="text"
                >JPG, PNG, PSD, AI, and SVG images up to 40MB with at least
                3000px on one side
              </span>
            </div>
          {/each}
        </div>
      </div>
    {:else}
      <EditInfo
        bind:title={fileName}
        bind:description
        bind:tags
        fileUrl={URL.createObjectURL(FileToUpload)}
        {submit}
        {cancel}
      />
    {/if}
  </div>

  {#if showNotif}
    <div
      role="alert"
      class="alert {isNotifError ? 'alert-error' : 'alert-success'}"
    >
      {#if isNotifError}
        <Icon icon="basil:info-triangle-outline" height="24" color="black"
        ></Icon>
      {:else}
        <Icon icon="basil:checked-box-outline" height="24" color="black"></Icon>
      {/if}
      <span
        >{isNotifError
          ? multiLang.uploadFailedMessage()
          : multiLang.uploadCompleteMessage()}</span
      >
    </div>
  {/if}
</div>

<style lang="scss">
  .page {
    max-width: 1830px;
    width: 100%;
    padding-top: 64px;
    padding-left: 16px;
    padding-right: 16px;
    margin-left: auto;
    margin-right: auto;
  }

  .container-1 {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    position: relative;

    h1 {
      font-size: 20px;
      font-weight: 800;
      margin-top: 40px;
      margin-bottom: 8px;
    }

    h2 {
      font-size: 14px;
      font-weight: 400;
      margin-bottom: 0;
      margin-top: 0;
    }
  }

  h2 {
    color: #656f79;
  }

  .container-2 {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 296px;
    background-color: #f7f7f7;
    border-radius: 8px;
    margin-top: 40px;
    margin-bottom: 40px;
    overflow: hidden;
    position: relative;
    animation: fadein--5BFpp 0.4s linear forwards;
    &::before {
      content: "";
      position: absolute;
      border: 4px dashed #ceced2;
      top: -3px;
      bottom: -3px;
      right: -3px;
      left: -3px;
      border-radius: 8px;
      pointer-events: none;
    }
  }

  .container-3 {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .mediaIconsContainer {
    display: flex;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 24px;
    position: relative;

    div {
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 100px;
      width: 64px;
      height: 64px;
      border: 1px solid #ebecf0;
      background-color: #fff;
      box-shadow: 0 2px 8px rgba(25, 27, 38, 0.06);

      .icon {
        font-style: normal;
        font-weight: 400;
        font-variant: normal;
        text-transform: none;
        font-size: 32px;
      }
    }
  }

  .videoIconContainer {
    transform: rotate(-15deg);
    left: -85%;
    top: 15%;
    z-index: 2;
    position: absolute;
  }

  .imageIconContainer {
    z-index: 1;
  }

  .musicIconContainer {
    transform: rotate(15deg);
    top: 15%;
    right: -85%;
    position: absolute;
  }

  .promptText {
    font-size: 16px;
    font-weight: 700;
    line-height: 24px;
    margin-bottom: 24px;
  }

  .infoSection {
    display: flex;
    align-items: center;
    margin-left: auto;
    margin-right: auto;
    font-size: 14px;
    line-height: 20px;
    color: #656f79;
  }

  .uploadsRemainingText {
    margin-right: 8px;
    display: flex;
    align-items: center;
  }

  .tooltiptext {
    visibility: hidden;
    top: 35%;
    left: 10%;
    right: 10%;
    background-color: rgb(12, 41, 59);
    color: #fff;
    text-align: center;
    padding: 5px 0;
    border-radius: 6px;
    padding: 4px;

    /* Position the tooltip text - see examples below! */
    position: absolute;
    z-index: 2;
  }

  .tooltipTriggerWrapper:hover .tooltiptext {
    visibility: visible;
  }

  .tooltipTriggerWrapper .tooltiptext::after {
    content: " ";
    position: absolute;
    top: 100%; /* At the bottom of the tooltip */
    left: 50%;
    margin-left: -5px;
    border-width: 5px;
    border-style: solid;
    border-color: black transparent transparent transparent;
  }

  .guidelinesContainer {
    width: 100%;
    height: 50%;
    -ms-grid-columns: none;
    grid-template-columns: none;
    grid-template-rows: repeat(auto-fit, minmax(40px, 1fr));
    row-gap: 16px;
    padding-bottom: 16px;

    div {
      justify-content: flex-start;
      display: flex;
      text-align: left;
      margin-bottom: 10px;

      .icon {
        width: 24px;
        font-size: 24px;
        line-height: 20px;
        margin-right: 20px;
        font-weight: 400;
      }

      .text {
        color: #656f79;
        font-size: 14px;
        line-height: 20px;
      }
    }

    @media screen and (min-width: 769px) {
      width: 936px;
      display: -ms-grid;
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(256px, 1fr));
      column-gap: 24px;
      row-gap: 24px;

      div {
        justify-content: space-between;
        margin-bottom: auto;
      }
    }
  }

  .zeroUploadsContainer {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-bottom: 96px;
  }
</style>
