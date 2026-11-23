export class ApiError extends Error {
    status: number;

    constructor(status: number, message: string) {
        super(message);
        this.status = status;
    }
}

// The Go side answers with {"error": "..."}, {"error": {"error": "..."}} or
// {"message": "..."}, so pick whichever one is there.
export function readMessage(body: unknown): string | null {
    if (!body || typeof body !== 'object')
        return null;

    const record = body as Record<string, unknown>;

    if (typeof record.error === 'string')
        return record.error;
    if (typeof record.message === 'string')
        return record.message;
    if (record.error && typeof record.error === 'object')
        return readMessage(record.error);

    return null;
}
