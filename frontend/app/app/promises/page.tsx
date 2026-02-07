"use client";

import { useEffect, useState } from "react";
import { getAllPromises } from "@/lib/promises";

type PromiseItem = {
    id: number;
    detail: string;
    status: string;
    announced_at: string;
    politician: { name: string; party: string };
};

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardAction,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

export default function PromisesPage() {
    const [items, setItems] = useState<PromiseItem[]>([]);
    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            alert("Please login first");
            window.location.href = "/login";
            return;
        }

        getAllPromises()
            .then(setItems)
            .catch(err => alert(err.message));

    }, []);
    console.log(items);
    return (
        <div className="w-full h-screen">
            <div className="flex flex-col justify-center items-center p-4 gap-4 flex-wrap">
                <h1>All Promises</h1>
                <div className="w-full grid grid-cols-5 gap-4">
                    {items.map(p => (
                        <Card key={p.id} className="w-full">
                            <CardHeader>
                                <CardTitle>{p.detail}</CardTitle>
                                <CardDescription>
                                    <Badge variant="secondary">
                                        {p.status}
                                    </Badge>
                                </CardDescription>
                            </CardHeader>
                            <CardFooter className="flex flex-col gap-2">
                                <div>
                                    Politician: {p.politician.name} ({p.politician.party})
                                </div>
                                <Button variant="link" className="w-full" asChild>
                                    <a href={`promises/${p.id}`}>View detail</a>
                                </Button>
                            </CardFooter>
                        </Card>
                    ))}
                </div>

            </div>
        </div>
    )
}
