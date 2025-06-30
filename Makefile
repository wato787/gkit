build:
	go build -o gkit .
	go build -o gs ./cmd/gs
	go build -o ga ./cmd/ga
	go build -o gc ./cmd/gc
	go build -o gp ./cmd/gp

install: build
	cp gkit /usr/local/bin/
	cp gs /usr/local/bin/
	cp ga /usr/local/bin/
	cp gc /usr/local/bin/
	cp gp /usr/local/bin/

install-completion:
	@echo "Installing shell completion..."
	@echo "For bash, add this to your ~/.bashrc or ~/.bash_profile:"
	@echo "source <(gs completion bash)"
	@echo "source <(ga completion bash)"
	@echo "source <(gc completion bash)"
	@echo "source <(gp completion bash)"
	@echo ""
	@echo "For zsh, add this to your ~/.zshrc:"
	@echo "source <(gs completion zsh)"
	@echo "source <(ga completion zsh)"
	@echo "source <(gc completion zsh)"
	@echo "source <(gp completion zsh)"

clean:
	rm -f gkit gs ga gc gp

.PHONY: build install install-completion clean