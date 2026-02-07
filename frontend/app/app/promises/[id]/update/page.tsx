"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { getPromiseDetail, addPromiseUpdate } from "@/lib/promises";

import { Button } from "@/components/ui/button"
import {
    Card,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"

export default function PromiseDetail() {
    const params = useParams();
    const router = useRouter();
    const id = Number(params.id);

    const [role, setRole] = useState<string | null>(null);
    const [detail, setDetail] = useState<any>(null);
    const [note, setNote] = useState("");
    const [error, setError] = useState("");

    useEffect(() => {
        setRole(localStorage.getItem("role"));
    }, []);

    console.log("id =", id);
    console.log("role =", role);

    useEffect(() => {
        if (!id) return;

        getPromiseDetail(id)
            .then(data => {
                if (data.status === "failed") {
                    alert("คำสัญญานี้ไม่สามารถอัปเดตได้");
                    router.push(`/promises/${id}`);
                    return;
                }
                setDetail(data);
            })
            .catch(() => router.push("/promises"));
    }, [id, router]);

    if (role === null) return <p>Loading...</p>;
    async function submit() {
        if (!note.trim()) {
            setError("กรุณากรอกรายละเอียด");
            return;
        }

        try {
            await addPromiseUpdate(id, note);
            router.push(`/promises/${id}`);
        } catch (err: any) {
            setError(err.message);
        }
    }

    if (!detail) return <p>Loading...</p>;

    return (
        <div className="w-full h-screen">
            <div className="flex flex-col justify-center items-center p-4 gap-4 flex-wrap">
                <h1>Promise Update</h1>
                <div className="w-3xl flex justify-center items-center gap-4">
                    <Card key={detail.id} className="w-full">
                        <CardHeader>
                            <CardTitle>{detail.detail}</CardTitle>
                            <CardDescription>
                                <Textarea
                                    rows={5}
                                    value={note}
                                    onChange={e => setNote(e.target.value)}
                                    placeholder="กรอกรายละเอียดการอัปเดตที่นี่"
                                />
                                {error && <p style={{ color: "red" }}>{error}</p>}
                            </CardDescription>
                        </CardHeader>
                        <CardFooter className="flex justify-center gap-2">
                            <Button className="bg-green-500 hover:bg-green-600 font-bold" onClick={submit}>บันทึก</Button>
                            <Button className="bg-red-500 hover:bg-red-600 font-bold" onClick={() => router.push(`/promises/${id}`)}>
                                ยกเลิก
                            </Button>
                        </CardFooter>
                    </Card>
                </div>
            </div>
        </div>
    )
}
