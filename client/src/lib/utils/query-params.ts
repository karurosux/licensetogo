import { goto } from '$app/navigation';

/**
 * This is a helper function to replace the current URL state with the given
 * query parameters. This is useful when you want to preserve the current
 * page state but change the query parameters.
 */
export const replaceStateWithQuery = (values: Record<string, string | number>) => {
	const searchParams = new URLSearchParams();
	for (let [k, v] of Object.entries(values)) {
		if (!!v) {
			searchParams.set(k, v.toString());
		} else {
			searchParams.delete(k);
		}
	}
	goto(Object.keys(values).length ? `?${searchParams.toString()}` : '');
};

/**
 * This is a helper function to merge the current URL state with the given
 * query parameters. This is useful when you want to preserve the current
 * page state but change the query parameters.
 */
export const mergeStateWithQuery = (values: Record<string, string | number>) => {
	const currentParams = getQueryParams();
	const mergedParams = { ...currentParams, ...values };
	replaceStateWithQuery(mergedParams);
};

/**
 * This is a helper function to get the current URL query parameters.
 */
export const getQueryParams = () => {
	const url = new URL(window.location.href);
	return Object.fromEntries(url.searchParams.entries());
};
