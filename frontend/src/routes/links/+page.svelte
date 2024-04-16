<script>
  import { createLink } from '$lib/repositories/links'
  import { urlByShortId, linkModel } from '$lib/helpers/link'

  async function shortenLink() {
    let response = await createLink(linkModel(url));
    let data = await response.json();

    if (response.status === 201) {
      shortenedUrl = urlByShortId(data.id);
    } else {
      console.log(response);
    }
  }

  let url;
  let shortenedUrl;
</script>

<div class="text-center">
  <h1 class="mb-2">Skratite link</h1>
  <div class="mb-4">
    <input bind:value={url} id="url" type="text" placeholder="Paste your link here" />
    <button on:click={shortenLink} class="btn btn-small btn-primary">Shorten</button>
  </div>
  {#if shortenedUrl}
    <div class="mb-4">
      <p>
        Your shortened link is:
        {shortenedUrl}
      </p>
    </div>
  {/if}
</div>