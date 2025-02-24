/**
 * Catches a promise and returns an object with the data and error
 * @param {Promise<any>} promise - The promise to catch
 * @returns {Promise<{ok: boolean, data: any, error: any}>}
 */
export const catchPromise = async (promise) => {
	try {
		const data = await promise;
		return {
			ok: true,
			data,
			error: null
		};
	} catch (error) {
		return {
			ok: false,
			data: null,
			error
		};
	}
};
