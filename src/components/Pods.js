import React, { useEffect, useState } from "react";
import axios from "axios";

function Pods() {
  const [pods, setPods] = useState([]);

  useEffect(() => {
    axios
      .get("/api/v1/namespaces/default/pods")
      .then((res) => setPods(res.data.items))
      .catch((err) => console.error("Error fetching pods:", err));
  }, []);

  return (
    <div>
      <h2 className="text-xl font-bold mb-2">Pods in 'default' namespace</h2>
      <ul className="list-disc pl-4">
        {pods.length === 0 ? (
          <li>No pods found</li>
        ) : (
          pods.map((pod) => (
            <li key={pod.metadata.name}>
              {pod.metadata.name} — {pod.status.phase}
            </li>
          ))
        )}
      </ul>
    </div>
  );
}

export default Pods;
