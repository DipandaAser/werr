<script lang="ts">
  export let src: string =
    "https://pixabay.com/static/img/profile_images/green.svg";
  export let alt: string = "Avatar";
  export let onClick: () => void;
</script>

<div class="container">
  <div class="triggerWrapper">
    <button on:click={onClick} class="userMenuTrigger" type="button">
      <span class="label"
        ><div class="container">
          <img {src} class="avatar" {alt} />
        </div></span
      ></button
    >
  </div>
</div>

<style lang="scss">
  .container {
    display: inline-flex;
    position: relative;
    overflow: visible;
    text-align: left;

    .triggerWrapper {
      display: flex;

      .userMenuTrigger {
        background: transparent;
        color: #fff;
        border: none;
        padding: 0;
        -webkit-appearance: button;
        appearance: button;
        border-radius: 24px;
        height: 40px;
        box-sizing: border-box;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        transition: all 0.1s ease-in;
        text-decoration: none;
        white-space: nowrap;
        font-weight: 600;
        font-size: 14px;
        overflow: visible;
        text-transform: none;
        margin: 0;

        .label {
          font-weight: 600;
          color: inherit;
          display: flex;
          align-items: center;
          justify-content: center;

          .container {
            width: 40px;
            height: 40px;
            overflow: hidden;
            border-radius: 50%;
            display: inline-flex;

            .avatar {
              width: 100%;
              height: 100%;
              display: inline-flex;
              align-items: center;
              justify-content: center;
              color: #656f79;
              object-fit: cover;
              background: #fff;
            }
          }
        }
      }
    }
  }
</style>
