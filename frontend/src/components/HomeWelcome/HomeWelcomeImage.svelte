<div class="base">
  <div class="image">
    <img
      src="https://cdn.pixabay.com/index/2023/11/24/01-58-59-34_1440x550.jpg"
      alt="background of a yard with a tree and a house"
    />
  </div>
</div>

<style lang="scss">
  .base {
    position: absolute;
    width: 100%;
    height: 40vh;
    z-index: 9;
    color: white;
  }
  .image {
    position: absolute;
    width: 100%;
    img {
      transform: translateX(-50%) translateY(-50%) scale(0.1);
    }
  }
</style>
