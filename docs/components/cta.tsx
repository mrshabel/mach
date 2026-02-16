"use client";

import { useState } from "react";

export function CTASection() {
    const [copied, setCopied] = useState(false);

    const copyToClipboard = () => {
        navigator.clipboard.writeText("go get github.com/mrshabel/mach");
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
    };

    return (
        <div className="relative py-32 bg-gradient-to-b from-slate-950 to-slate-900">
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(59,130,246,0.05),transparent_70%)]"></div>

            <div className="relative max-w-4xl mx-auto px-6 text-center">
                <h2 className="text-4xl lg:text-5xl font-bold text-white mb-6">
                    Ready to build something?
                </h2>
                <p className="text-xl text-slate-400 mb-10">
                    Stop configuring. Start coding.
                </p>

                <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
                    <div className="flex items-center gap-3 px-6 py-3 bg-slate-800/50 rounded-lg border border-slate-700 font-mono text-sm text-slate-300">
                        <span className="text-slate-500">$</span>
                        <code>go get github.com/mrshabel/mach</code>
                        <button
                            onClick={copyToClipboard}
                            className="ml-2 text-blue-400 hover:text-blue-300 transition-colors"
                        >
                            {copied ? (
                                <svg
                                    className="w-5 h-5"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M5 13l4 4L19 7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    className="w-5 h-5"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                                    />
                                </svg>
                            )}
                        </button>
                    </div>
                </div>

                <div className="mt-12 text-sm text-slate-500">
                    MIT Licensed • Built by{" "}
                    <a
                        href="https://github.com/mrshabel"
                        className="text-blue-400 hover:text-blue-300 transition-colors"
                    >
                        @mrshabel
                    </a>
                </div>
            </div>
        </div>
    );
}
