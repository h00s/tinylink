<script>
  import { createLink } from '$lib/repositories/links'
  import { urlByShortId, linkModel } from '$lib/helpers/link'

  async function shortenLink() {
    let response = await createLink(linkModel(url));
    let data = await response.json();

    if (response.status === 201) {
      shortenedUrl = urlByShortId(data.id);
    } else {
      if (data.description) {
        alert(data.description);
      } else {
        alert('An error occurred while shortening the link');
      }
      // console.log(response);
    }
  }

  let url;
  let shortenedUrl;
</script>

<div class="text-center">
  <div class="flex flex-col md:flex-row">
    <div class="flex-1 md:mr-4">
      <label class="input input-bordered flex items-center gap-2">
        🔗
        <input bind:value={url} id="url" type="text" class="w-full text-dark-blue" placeholder="Zalijepite link koji želite skratiti..." />
      </label>
    </div>
    <div class="mt-4 md:mt-0">
      <button on:click={shortenLink} class="btn btn-small btn-primary w-full md:w-auto">Skrati!</button>
    </div>
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