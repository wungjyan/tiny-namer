export namespace file {
	
	export class FileInfo {
	    fullName: string;
	    path: string;
	    dir: string;
	    ext: string;
	    baseName: string;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fullName = source["fullName"];
	        this.path = source["path"];
	        this.dir = source["dir"];
	        this.ext = source["ext"];
	        this.baseName = source["baseName"];
	    }
	}
	export class SelectFilesResult {
	    success: boolean;
	    files: FileInfo[];
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new SelectFilesResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.files = this.convertValues(source["files"], FileInfo);
	        this.message = source["message"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace highlight {
	
	export class HighlightRequest {
	    fileName: string;
	    baseName: string;
	    pattern: string;
	    ignoreExt: boolean;
	    ignoreCase: boolean;
	    global: boolean;
	
	    static createFrom(source: any = {}) {
	        return new HighlightRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.baseName = source["baseName"];
	        this.pattern = source["pattern"];
	        this.ignoreExt = source["ignoreExt"];
	        this.ignoreCase = source["ignoreCase"];
	        this.global = source["global"];
	    }
	}

}

export namespace regex {
	
	export class RenameRequest {
	    fileInfo: file.FileInfo;
	    pattern: string;
	    replacement: string;
	    ignoreExt: boolean;
	    ignoreCase: boolean;
	    global: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RenameRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileInfo = this.convertValues(source["fileInfo"], file.FileInfo);
	        this.pattern = source["pattern"];
	        this.replacement = source["replacement"];
	        this.ignoreExt = source["ignoreExt"];
	        this.ignoreCase = source["ignoreCase"];
	        this.global = source["global"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RenameResult {
	    success: boolean;
	    fullName: string;
	    path: string;
	    dir: string;
	    ext: string;
	    baseName: string;
	    newFullName: string;
	    newPath: string;
	    newDir: string;
	    newExt: string;
	    newBaseName: string;
	    message: string;
	    previewOnly: boolean;
	    changed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RenameResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.fullName = source["fullName"];
	        this.path = source["path"];
	        this.dir = source["dir"];
	        this.ext = source["ext"];
	        this.baseName = source["baseName"];
	        this.newFullName = source["newFullName"];
	        this.newPath = source["newPath"];
	        this.newDir = source["newDir"];
	        this.newExt = source["newExt"];
	        this.newBaseName = source["newBaseName"];
	        this.message = source["message"];
	        this.previewOnly = source["previewOnly"];
	        this.changed = source["changed"];
	    }
	}
	export class ValidateResult {
	    valid: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ValidateResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.valid = source["valid"];
	        this.message = source["message"];
	    }
	}

}

