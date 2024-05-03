<script>
  import FlexContainer from '$lib/components/FlexContainer.svelte';
  import Authorize from '$lib/components/Authorize.svelte';
  import Footer from '$lib/components/Footer.svelte';

  export let data;
</script>


<FlexContainer>
  <h1 class="pt-16 pb-2 font-extrabold text-center">
    {data.url}
  </h1>
  <Authorize link={data.link} />
</FlexContainer>

<Footer />