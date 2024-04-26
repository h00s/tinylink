<script>
  import Footer from '$lib/components/Footer.svelte';

  export let data;
</script>

<div class="flex justify-center">
  <div class="w-4/5 lg:w-1/2">
    <h1 class="pt-16 pb-2 font-extrabold text-center">
      {data.url}
    </h1>
  </div>
</div>

<Footer />