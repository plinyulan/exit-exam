"use client";

import { useEffect, useState } from "react";
import {
  getPoliticians,
  getPromisesByPolitician,
} from "@/lib/politicians";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

type Politician = {
  id: number;
  politician_code: string;
  name: string;
  party: string;
};

type PromiseItem = {
  id: number;
  detail: string;
  status: string;
};

export default function PoliticiansPage() {
  const [politicians, setPoliticians] = useState<Politician[]>([]);
  const [selected, setSelected] = useState<Politician | null>(null);
  const [promises, setPromises] = useState<PromiseItem[]>([]);
  const [loading, setLoading] = useState(false);

  // โหลดนักการเมือง
  useEffect(() => {
    getPoliticians()
      .then(setPoliticians)
      .catch(err => alert(err.message));
  }, []);

  async function selectPolitician(p: Politician) {
    setSelected(p);
    setLoading(true);
    setPromises([]);
    try {
      const data = await getPromisesByPolitician(p.id);
      setPromises(data);
    } catch (err: any) {
      alert(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">นักการเมือง & คำสัญญา</h1>

      <div className="grid grid-cols-12 gap-6">
        {/* LEFT: Politicians */}
        <div className="col-span-12 md:col-span-4 space-y-4">
          {politicians.map(p => (
            <Card
              key={p.id}
              onClick={() => selectPolitician(p)}
              className={`cursor-pointer transition hover:shadow-md ${
                selected?.id === p.id ? "border-primary" : ""
              }`}
            >
              <CardHeader>
                <CardTitle className="text-lg">{p.name}</CardTitle>
                <CardDescription>{p.party}</CardDescription>
              </CardHeader>
            </Card>
          ))}
        </div>

        {/* RIGHT: Promises */}
        <div className="col-span-12 md:col-span-8 space-y-4">
          {!selected && (
            <Card className="h-full flex items-center justify-center">
              <CardContent className="text-muted-foreground">
                เลือกนักการเมืองเพื่อดูคำสัญญา
              </CardContent>
            </Card>
          )}

          {selected && (
            <>
              <h2 className="text-xl font-semibold">
                คำสัญญาของ {selected.name}
              </h2>

              {!loading && promises.length === 0 && (
                <Card>
                  <CardContent className="text-muted-foreground">
                    ไม่มีคำสัญญา
                  </CardContent>
                </Card>
              )}

              {!loading &&
                promises.map(pr => (
                  <Card key={pr.id}>
                    <CardHeader>
                      <CardTitle className="text-base">
                        {pr.detail}
                      </CardTitle>
                    </CardHeader>
                    <CardFooter>
                      <Badge variant="secondary">{pr.status}</Badge>
                    </CardFooter>
                  </Card>
                ))}
            </>
          )}
        </div>
      </div>
    </div>
  );
}
