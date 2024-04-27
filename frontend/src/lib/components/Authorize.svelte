<script>
  import { authLink } from '$lib/repositories/links';

  let error;
</script>

<div class="text-center">
  <div class="flex flex-col md:flex-row">
    <div class="flex-1 md:mr-4">
      <label class="input input-bordered flex items-center gap-2">
        🔒
        <input id="url" type="password" class="w-full text-dark-blue" />
      </label>
    </div>
    <div class="mt-4 md:mt-0">
      <button on:click={authLink} class="btn btn-small btn-primary w-full md:w-auto">Unesi</button>
    </div>
  </div>
  {#if error}
    <div class="mb-4">
      <p>
        Your shortened link is:
        {error}
      </p>
    </div>
  {/if}
</div>