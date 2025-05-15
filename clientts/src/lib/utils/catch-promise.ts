export type CatchResponse<T = any> = {
	ok: boolean;
	data: T | null;
	error: Error | null;
};

export const catchPromise = async <T>(promise: Promise<T>): Promise<CatchResponse<T>> => {
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
			error: error as Error
		};
	}
};
