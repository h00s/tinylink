import { urlByShortId } from '$lib/helpers/link';

export async function load({ fetch, params }) {
  const url = urlByShortId(params.link);

  return {
    url: url,
  }
}