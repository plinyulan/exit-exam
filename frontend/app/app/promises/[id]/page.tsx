"use client";

import { useEffect, useState } from "react";
import { isAdmin } from "@/lib/auth";
import { useParams } from "next/navigation";
import { getPromiseDetail } from "@/lib/promises";
import {
    Card,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

export default function PromiseDetail() {
    const params = useParams();
    const id = Number(params.id);

    const [data, setData] = useState<any>(null);

    useEffect(() => {
        if (!id) return;

        getPromiseDetail(id)
            .then(setData)
            .catch(err => alert(err.message));
    }, [id]);

    if (!data) return <p>Loading...</p>;
    return (
        <div className="w-full h-screen">
            <div className="flex flex-col justify-center items-center p-4 gap-4 flex-wrap">
                <h1>Promise Detail</h1>
                <div className="w-3xl flex justify-center items-center gap-4">
                    <Card key={data.id} className="w-full">
                        <CardHeader>
                            <CardTitle>{data.detail}</CardTitle>
                            <CardDescription>
                                {/* <Badge variant="secondary">
                                    {data.status}
                                </Badge> */}
                                {data.updates.map((u: any) => (
                                    <div key={u.id}>
                                        {new Date(u.updated_at).toLocaleString()} - {u.note}
                                    </div>
                                ))}
                            </CardDescription>
                        </CardHeader>
                        <CardFooter className="flex flex-col gap-2">
                            <div>
                                Politician: {data.politician.name} ({data.politician.party})
                            </div>
                            {isAdmin() && data.status !== "failed" && (
                               
                                    <a className="border-0 rounded-2xl p-2 bg-green-500 text-white text-md font-bold hover:bg-green-600" href={`/promises/${params.id}/update`}>Update</a>

                            )}
                        </CardFooter>
                    </Card>
                </div>

            </div>
        </div>
    )
}
