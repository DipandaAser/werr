<script lang="ts">
  import Icon from "@iconify/svelte";
  let searchQuery = "";
  let searchCategorie = "images";
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
      <h1>Stunning media library for your next video content</h1>
      <h2>
        Over 4.2 thousand+ high quality images, videos and audio shared by our
        amazing community for your content creation journey.
      </h2>
      <div class="search-box">
        <form class="search-form" action="#" method="get">
          <!-- TODO: to be replaced by form -->
          <button
            class="submit"
            type="submit"
            aria-label="Search for {searchCategorie}"
          >
            <Icon icon="basil:search-solid" color="black" />
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
            <button class="dropdown" type="button">
              <span class="label"
                >{searchCategorie.charAt(0).toUpperCase() +
                  searchCategorie.slice(1)}</span
              >
              <Icon
                class="icon"
                icon="basil:caret-down-solid"
                height="24"
                width="24"
              />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="scss">
  $hoverHeaderColor: rgba(182, 175, 175, 0.151);
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
        padding: 0 0 0 24px;
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
              color: #656f79;
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
              color: inherit;
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
      }
    }

    // DEBUG
    // background-color: aqua;
    //border-bottom: 1px solid salmon;
  }
</style>
