---

# Mini Headlamp 🚦

A lightweight React-based UI to visualize Kubernetes resources from your local cluster (created using `kind`). This is a simple prototype inspired by the Kubernetes Headlamp project.

---

## Features

* Connects to a Kubernetes cluster running locally (via `kind`)
* Displays pods in the `default` namespace with their status
* Fetches data through `kubectl proxy` to avoid CORS issues
* Easy to run and customize for Kubernetes resource visualization

---

## Prerequisites

* [Node.js](https://nodejs.org/) (v16 or later recommended)
* [kubectl](https://kubernetes.io/docs/tasks/tools/) CLI installed and configured
* [kind](https://kind.sigs.k8s.io/) to create a local Kubernetes cluster
* Running Kubernetes cluster (created by `kind`)
* `kubectl proxy` running on your machine

---

## Setup and Run

1. **Create a Kubernetes cluster using kind**

```bash
kind create cluster
```

2. **Start the Kubernetes API proxy**

```bash
kubectl proxy
```

This will serve the API at `http://localhost:8001`.

3. **Clone this repository**

```bash
git clone <your-repo-url>
cd mini-headlamp
```

4. **Install dependencies**

```bash
npm install
```

5. **Start the React app**

```bash
npm start
```

6. **Open your browser**

Visit [http://localhost:3001](http://localhost:3001) to view Mini Headlamp.

---

## Important Notes

* The React app uses the `"proxy": "http://localhost:8001"` setting in `package.json` to forward API calls to the Kubernetes proxy and avoid CORS issues.
* Ensure your Kubernetes cluster has pods running (e.g., deploy an nginx pod) to see meaningful data.
* This is a development prototype; the production build is not yet optimized.

---

## Deploy a test pod (optional)

To deploy an example pod and see it in Mini Headlamp:

```bash
kubectl run nginx --image=nginx --restart=Never
```

---

## Troubleshooting

* If you don’t see pods listed, verify:

  * Your cluster is running (`kubectl get nodes`)
  * `kubectl proxy` is running on `localhost:8001`
  * Pods exist in the `default` namespace (`kubectl get pods`)
  * The React app has restarted after adding `"proxy"` in `package.json`

---

## Future Improvements

* Support more Kubernetes resources (Deployments, Services, ConfigMaps)
* Add pod logs and metrics viewing
* Authentication support for remote clusters
* UI/UX enhancements for a smoother experience

---

## License

MIT License

---

