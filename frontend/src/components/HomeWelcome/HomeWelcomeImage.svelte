<script lang="ts">
  import { CategoriesList } from "$lib/models/categorie";
  import { searchStore } from "../../stores/searchStore";
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";
  let searchQuery = "";
  let searchCategorie = "images";
  let cursorInSearchBox = false;
  let catChooserDropdownIsOpen = false;
  export let welcomeTitle: string;
  export let description: string;

  const selectedColor = "#00ab6b";
  const unSelectedColor = "#656f79";

  onMount(() => {
    const searchInput = document.querySelector(".search-box") as HTMLElement;
    searchInput.addEventListener("click", () => {
      cursorInSearchBox = true;
    });

    searchInput.addEventListener("blur", () => {
      cursorInSearchBox = false;
    });
  });

  function toggleContainer() {
    catChooserDropdownIsOpen = !catChooserDropdownIsOpen;
    if (catChooserDropdownIsOpen) {
      const dropdownMenuContainer = document.querySelector(
        ".dropdownMenuContainer"
      ) as HTMLElement;
      dropdownMenuContainer.style.display = "block";
    } else {
      const dropdownMenuContainer = document.querySelector(
        ".dropdownMenuContainer"
      ) as HTMLElement;
      dropdownMenuContainer.style.display = "none";
    }
  }
</script>

<div class="base">
  <div class="image">
    <picture>
      <source
        srcset="https://cdn.pixabay.com/index/2023/11/24/01-58-59-34_1920x550.jpg"
        media="(min-width: 1440px)"
      />
      <source
        srcset="https://cdn.pixabay.com/index/2023/11/24/01-58-59-34_640.jpg"
        media="(max-width: 640px)"
      />
      <img
        src="https://cdn.pixabay.com/index/2023/11/24/01-58-59-34_1440x550.jpg"
        class="imageBanner"
        alt="Free mixed hero backgrounds"
      />
    </picture>
  </div>
  <!-- This is just the blur filter-->
  <div class="overlay"></div>
  <div class="welcomeMessageBox">
    <div class="wrapper">
      <h1>{welcomeTitle}</h1>
      <h2>{description}</h2>
      <div class="search-box">
        <form class="search-form" action="#" method="get">
          <!-- TODO: to be replaced by form -->
          <button
            class="submit"
            type="submit"
            aria-label="Search for {searchCategorie}"
          >
            <Icon
              icon="basil:search-solid"
              color={cursorInSearchBox ? "green" : "black"}
            />
          </button>
          <input
            type="search"
            name="search"
            placeholder="Search for {searchCategorie}"
            bind:value={searchQuery}
          />
        </form>
        <div class="chooseCat">
          <div class="chooseCatContainer">
            <button class="dropdown" type="button" on:click={toggleContainer}>
              <span class="label">{$searchStore.categorie}</span>
              <Icon
                class="icon"
                icon={catChooserDropdownIsOpen
                  ? "basil:caret-up-solid"
                  : "basil:caret-down-solid"}
                color={catChooserDropdownIsOpen ? "#00ab6b" : "black"}
                height="24"
                width="24"
              />
            </button>
          </div>

          <div class="dropdownMenuContainer">
            <div class="dropdownMenu">
              {#each CategoriesList as categorie}
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div
                  aria-label="Choose {categorie.label}"
                  id="cat-{categorie.value}"
                  class="dropdownMenuItem"
                  on:click={() => {
                    $searchStore.categorie = categorie.label;
                    toggleContainer();
                  }}
                >
                  <div class="icon">
                    <Icon
                      class="icon-svg"
                      icon={categorie.icon}
                      height="24"
                      width="24"
                      color={categorie.label === $searchStore.categorie
                        ? selectedColor
                        : unSelectedColor}
                    />
                  </div>
                  <!-- svelte-ignore a11y-label-has-associated-control -->
                  <label
                    style="color: {categorie.label === $searchStore.categorie
                      ? selectedColor
                      : unSelectedColor};">{categorie.label}</label
                  >
                </div>
              {/each}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="scss">
  $hoverHeaderColor: rgba(182, 175, 175, 0.151);
  $chooseCatTextColor: #656f79;
  .base {
    min-height: 464px;
    position: relative;
    width: 100%;
    color: white;
    z-index: 9;
    font-size: 14px;

    // DEBUG
    //background-color: aqua;
    //border-bottom: 1px solid salmon;
  }

  .image {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .imageBanner {
    position: absolute;
    min-width: 1000%;
    min-height: 1000%;
    left: 50%;
    top: 50%;
    transform: translateX(-50%) translateY(-50%) scale(0.1);
  }

  .overlay {
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(
      0deg,
      rgba(25, 27, 38, 0.32),
      rgba(25, 27, 38, 0.72)
    );
  }

  .welcomeMessageBox {
    padding: 0 16px;
    min-height: 464px;
    position: relative;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    text-align: left;

    .wrapper {
      width: 100%;
      max-width: 840px;
      display: flex;
      flex-direction: column;
      align-items: flex-start;

      h1 {
        font-size: 24px;
        margin: 0 0 8px;
        font-weight: 800;
      }

      h2 {
        font-weight: 400;
        font-size: 14px;
        margin: 0 0 24px;
        color: #fff;
      }

      .search-box {
        background-color: #fff;
        border: 1px solid #f7f7f7;
        height: 56px;
        padding: 0 8px 0 24px;
        border-radius: 56px;
        box-shadow: 0 16px 40px rgba(25, 27, 38, 0.08);
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        position: relative;

        .search-form {
          flex: auto;
          display: flex;
          align-items: center;
          justify-content: center;
          min-width: 32px;

          .submit {
            border-style: none;
            padding: 0;
            color: #656f79;
            display: flex;
            align-items: center;
            justify-content: center;
            border: none;
            background: none;
            cursor: pointer;
            font-size: 24px;
            font-weight: 400;
            overflow: visible;
            text-transform: none;
            margin: 0;

            span {
              font-style: normal;
              font-weight: 400;
              font-variant: normal;
              text-transform: none;
            }
          }

          input {
            color: #656f79;
            opacity: 1;
            color: #656f79;
            flex: auto;
            padding: 8px;
            border: none;
            background: none;
            margin: 0;
            min-width: 30px;
            font-size: 14px;
            font-weight: 400;
            overflow: visible;
            outline: none;
          }
        }

        .chooseCat {
          display: inline-flex;
          position: relative;
          overflow: visible;
          text-align: left;

          .chooseCatContainer {
            display: flex;
            .dropdown {
              margin-right: 8px;
              border-radius: 56px;
              background: transparent;
              color: $chooseCatTextColor;
              border: none;
              height: 40px;
              padding: 0 16px;
              display: inline-flex;
              align-items: center;
              justify-content: start;
              cursor: pointer;
              transition: all 0.1s ease-in;
              text-decoration: none;
              white-space: nowrap;
              font-weight: 600;
              font-size: 14px;
              overflow: visible;
              text-transform: none;
              margin: 0;

              &:hover {
                background-color: $hoverHeaderColor;
              }
            }

            .label {
              font-weight: 600;
              display: flex;
              align-items: center;
              justify-content: center;
              cursor: pointer;
            }

            .icon {
              font-size: 20px;
              margin-left: 8px;
              cursor: pointer;
              font-style: normal;
              font-weight: 400;
              font-variant: normal;
              text-transform: none;
            }
          }
        }

        .dropdownMenuContainer {
          display: none;
          color: $chooseCatTextColor;
          background: #fff;
          top: calc(100% + 4px);
          right: 0;
          position: absolute;
          min-width: 216px;
          border-radius: 8px;
          box-shadow:
            0 0 0 1px rgba(14, 19, 24, 0.07),
            0 2px 18px rgba(14, 19, 24, 0.25);
          z-index: 90;

          .dropdownMenu {
            display: flex;
            flex-direction: column;
            padding: 12px;

            .dropdownMenuItem {
              // color: #656f79;
              display: flex;
              align-items: center;
              padding: 8px 12px;
              text-decoration: none;
              flex-wrap: nowrap;
              white-space: nowrap;
              cursor: pointer;
              border-radius: 8px;
              user-select: none;

              &:hover {
                background-color: $hoverHeaderColor;
                color: black;
              }

              .icon {
                font-size: 20px;
                flex: initial;
                display: flex;
                align-items: center;
                justify-content: center;
                min-width: 24px;
                margin-right: 8px;

                .icon-svg {
                  font-style: normal;
                  font-weight: 400;
                  font-variant: normal;
                  text-transform: none;
                }
              }

              label {
                white-space: nowrap;
                cursor: pointer;
                font-size: 14px;
                font-weight: 400;
              }
            }
          }
        }
      }
    }

    // DEBUG
    // background-color: aqua;
    //border-bottom: 1px solid salmon;
  }
</style>
