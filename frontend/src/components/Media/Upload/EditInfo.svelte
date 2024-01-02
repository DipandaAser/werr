<script lang="ts">
  import Icon from "@iconify/svelte";
  import "../../../app.css";
  import * as multiLang from "$paraglide/messages";

  export let fileUrl: string;
  export let fileType: string = "";
  export let description: String = "";
  export let tags: String[] = [];
  export let title: String = "";
  export let submit: () => void;
  export let cancel: () => void;

  let tagsString: String = "";

  let userAlreadyEnterInTagInput: boolean = false;
  let showErrorForTagInput: boolean = false;
  let showErrorForTitleInput: boolean = false;
  let userAlreadyEnterInTitleInput: boolean = false;
  let userIsInTitleInput: boolean = false;

  let userIsInTagInput: boolean = false;

  $: showErrorForTagInput = tags.length < 3 && userAlreadyEnterInTagInput;
  $: showErrorForTitleInput =
    title.length === 0 && userAlreadyEnterInTitleInput;
</script>

<div id="card" style="border: 1px solid #ebecf0" class="card mt-7 mb-10">
  <figure>
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img src={fileUrl} alt="The media to upload" />
  </figure>
  <div class="card-body">
    <!-- TITLE -->
    <div class="label">
      <label for="title" class="label-text">{multiLang.title()}</label>
    </div>
    <input
      id="title"
      type="text"
      placeholder={multiLang.editMedia_TitlePlaceHolder()}
      class="input {showErrorForTitleInput
        ? userIsInTitleInput
          ? ''
          : 'input-error'
        : ''} input-bordered w-full"
      on:change={() => (userAlreadyEnterInTitleInput = true)}
      on:input={() => (userAlreadyEnterInTitleInput = true)}
      on:click={() => (userAlreadyEnterInTitleInput = true)}
      on:focus={() => {
        userAlreadyEnterInTitleInput = true;
        userIsInTitleInput = true;
      }}
      on:blur={() => (userIsInTitleInput = false)}
      bind:value={title}
    />
    <!-- TAGS -->
    <div class="label">
      <span class="label-text">{multiLang.tags()}</span>
    </div>
    <input
      id="tags"
      type="text"
      placeholder={multiLang.editMedia_TagsPlaceHolder()}
      class="input {showErrorForTagInput
        ? userIsInTagInput
          ? ''
          : 'input-error'
        : ''} input-bordered w-full"
      on:change={() => (userAlreadyEnterInTagInput = true)}
      on:input={() => (userAlreadyEnterInTagInput = true)}
      on:click={() => (userAlreadyEnterInTagInput = true)}
      on:focus={() => {
        userAlreadyEnterInTagInput = true;
        userIsInTagInput = true;
      }}
      on:blur={() => (userIsInTagInput = false)}
      on:keyup={(event) => {
        if (event.key === "Enter" || event.key === ",") {
          tagsString = tagsString.trim();
          //if the user enter a comma, we remove it
          if (tagsString[tagsString.length - 1] === ",") {
            tagsString = tagsString.slice(0, -1);
          }
          if (tagsString.length === 0) {
            return;
          }

          if (tagsString === ",") {
            tagsString = "";
            return;
          }
          //search if the tag exist in the array
          const tagExist = tags.find(
            (tag) => tag.toLowerCase() === tagsString.toLowerCase()
          );
          if (tagExist) {
            //if the tag exist, we don't add it
            tagsString = "";
            return;
          }
          //create a new tags array by combining the old one with the new one
          tags = [...tags, tagsString];
          tagsString = "";
        }
      }}
      bind:value={tagsString}
    />

    <div>
      {#each tags as tag, index}
        <span class="badge badge-primary badge-outline">
          {tag} &nbsp;
          <div class="tooltip" data-tip="Delete this tag">
            <button
              style="cursor: pointer;"
              on:click={() => {
                tags = tags.filter((_, i) => i !== index);
                console.log("tag deleted");
              }}
            >
              <Icon
                style="cursor: pointer;"
                icon="basil:trash-solid"
                color="red"
                width="12"
                on:click={() => {
                  tags = tags.filter((_, i) => i !== index);
                  console.log("tag deleted");
                }}
              />
            </button>
          </div>
        </span>
      {/each}
    </div>

    <!-- DESCRIPTION -->
    <div class="label">
      <label for="description" class="label-text"
        >{multiLang.description()}</label
      >
      <span class="label-text">{description.length}/300</span>
    </div>
    <textarea
      id="description"
      placeholder={multiLang.editMedia_DescriptionPlaceHolder()}
      class="textarea textarea-bordered w-full"
      maxlength="300"
      bind:value={description}
    ></textarea>
    <div class="card-actions justify-end">
      <button
        style="background-color: #ceced2;"
        class="btn rounded-full"
        on:click={cancel}>{multiLang.cancel()}</button
      >
      <button
        class="btn btn-primary rounded-full"
        on:click={() => {
          if (title.length === 0) {
            userAlreadyEnterInTitleInput = true;
            userIsInTitleInput = false;
            return;
          }
          if (tags.length < 3) {
            userAlreadyEnterInTagInput = true;
            userIsInTagInput = false;
            return;
          }
          submit();
        }}>{multiLang.submit()}</button
      >
    </div>
  </div>
</div>

<br />

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.white);
  }

  /* for screen with minimal width of 768 apply the class card-side to the first div */
  @media (min-width: 769px) {
    #card {
      @apply card-side;
      width: 968px;
    }

    figure {
      width: 360px;
    }
  }
</style>
