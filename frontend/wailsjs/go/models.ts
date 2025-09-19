export namespace main {
	
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
	export class RegexValidationResult {
	    valid: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new RegexValidationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.valid = source["valid"];
	        this.message = source["message"];
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

