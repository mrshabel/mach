import Link from "next/link";

export function HeroSection() {
    return (
        <div className="relative min-h-screen flex items-center">
            {/* Background Gradient */}
            <div className="absolute inset-0 bg-gradient-to-br from-blue-950 via-slate-900 to-slate-950"></div>
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_30%_20%,rgba(59,130,246,0.1),transparent_50%)]"></div>
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_70%_80%,rgba(14,165,233,0.08),transparent_50%)]"></div>

            <div className="relative max-w-7xl mx-auto px-6 lg:px-8 py-20 lg:py-32 grid lg:grid-cols-2 gap-16 items-center">
                {/* Left: Text Content */}
                <div className="space-y-8">
                    <div className="inline-flex items-center gap-2 px-3 py-1.5 rounded-full border border-blue-500/30 bg-blue-500/10 text-blue-400 text-sm font-medium">
                        <span className="relative flex h-2 w-2">
                            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-blue-400 opacity-75"></span>
                            <span className="relative inline-flex rounded-full h-2 w-2 bg-blue-500"></span>
                        </span>
                        Go 1.22+ Native Routing
                    </div>

                    <h1 className="text-5xl lg:text-7xl font-bold tracking-tight">
                        <span className="text-white">Stop fighting</span>
                        <br />
                        <span className="text-white">your framework.</span>
                        <br />
                        <span className="bg-gradient-to-r from-blue-400 via-cyan-400 to-teal-400 bg-clip-text text-transparent">
                            Start shipping.
                        </span>
                    </h1>

                    <p className="text-lg lg:text-xl text-slate-400 max-w-xl leading-relaxed">
                        Mach is a Go web framework that gets out of your way.
                        Zero deps, stdlib routing, context pooling. Built by
                        someone tired of framework bloat.
                    </p>

                    <div className="flex flex-wrap gap-4">
                        <Link
                            href="/docs"
                            className="group relative inline-flex items-center gap-2 px-6 py-3 bg-blue-600 hover:bg-blue-500 text-white font-semibold rounded-lg transition-all shadow-lg shadow-blue-500/20 hover:shadow-blue-500/40"
                        >
                            Read the docs
                            <svg
                                className="w-4 h-4 group-hover:translate-x-1 transition-transform"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor"
                            >
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth={2}
                                    d="M9 5l7 7-7 7"
                                />
                            </svg>
                        </Link>
                        <a
                            href="https://github.com/mrshabel/mach"
                            target="_blank"
                            rel="noopener noreferrer"
                            className="inline-flex items-center gap-2 px-6 py-3 bg-slate-800 hover:bg-slate-700 text-slate-200 font-semibold rounded-lg transition-all border border-slate-700"
                        >
                            <svg
                                className="w-5 h-5"
                                fill="currentColor"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    fillRule="evenodd"
                                    d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                                    clipRule="evenodd"
                                />
                            </svg>
                            GitHub
                        </a>
                    </div>

                    {/* Quick Stats */}
                    <div className="flex gap-8 pt-8 border-t border-slate-800">
                        <div>
                            <div className="text-3xl font-bold text-white">
                                0
                            </div>
                            <div className="text-sm text-slate-500">
                                Dependencies
                            </div>
                        </div>
                        <div>
                            <div className="text-3xl font-bold text-white">
                                ~
                            </div>
                            <div className="text-sm text-slate-500">
                                Overhead vs stdlib
                            </div>
                        </div>
                        <div>
                            <div className="text-3xl font-bold text-white">
                                100%
                            </div>
                            <div className="text-sm text-slate-500">
                                Stdlib compatible
                            </div>
                        </div>
                    </div>
                </div>

                {/* Right: Code Example */}
                <div className="relative">
                    {/* Floating decoration */}
                    <div className="absolute -top-4 -right-4 w-72 h-72 bg-blue-500/10 rounded-full blur-3xl"></div>

                    <div className="relative bg-slate-900/80 backdrop-blur-sm rounded-2xl border border-slate-800 shadow-2xl overflow-hidden">
                        {/* Terminal Header */}
                        <div className="flex items-center gap-2 px-4 py-3 bg-slate-950/50 border-b border-slate-800">
                            <div className="flex gap-1.5">
                                <div className="w-3 h-3 rounded-full bg-red-500/80"></div>
                                <div className="w-3 h-3 rounded-full bg-yellow-500/80"></div>
                                <div className="w-3 h-3 rounded-full bg-green-500/80"></div>
                            </div>
                            <span className="text-xs text-slate-500 ml-2">
                                main.go
                            </span>
                        </div>

                        {/* Code */}
                        <pre className="p-6 text-sm font-mono overflow-x-auto">
                            <code className="text-slate-300">
                                {`package main

import "github.com/mrshabel/mach"

func main() {
    app := mach.Default()
    
    app.GET("/", func(c *mach.Context) {
        c.JSON(200, map[string]string{
            "message": "blazingly fast",
        })
    })
    
    app.Run(":8080")
}`}
                            </code>
                        </pre>

                        {/* Terminal Output Simulation */}
                        <div className="px-6 pb-6 pt-2 border-t border-slate-800/50">
                            <div className="text-xs font-mono space-y-1">
                                <div className="text-slate-500">
                                    $ go run main.go
                                </div>
                                <div className="text-green-400">
                                    [mach] Server started on :8080
                                </div>
                                <div className="text-slate-600 flex items-center gap-2">
                                    <span className="inline-block w-1 h-3 bg-slate-400 animate-pulse"></span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
